package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"danglingmind.com/ddd/application"
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/service"
	"danglingmind.com/ddd/infrastructure/auth"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Blog struct {
	blogApp    application.BlogAppInterface
	userApp    application.UserAppInterface
	tagApp     application.TagAppInterface
	blogTagApp application.BlogTagAppInterface
	tagService service.TagServiceInterface
	tk         auth.TokenInterface
	au         auth.AuthInterface
}

type blogRequestPayload struct {
	Blog entity.Blog  `json:"blog"`
	Tags []entity.Tag `json:"tags"`
}

type blogResponse struct {
	entity.Blog
	Tags []entity.Tag
}

func NewBlog(
	b application.BlogAppInterface,
	u application.UserAppInterface,
	ta application.TagAppInterface,
	bta application.BlogTagAppInterface,
	ts service.TagServiceInterface,
	t auth.TokenInterface,
	a auth.AuthInterface) *Blog {
	return &Blog{
		blogApp:    b,
		userApp:    u,
		tagApp:     ta,
		blogTagApp: bta,
		tagService: ts,
		tk:         t,
		au:         a,
	}
}

// user auth required
func (bg *Blog) Save(w http.ResponseWriter, r *http.Request) {
	// validate user
	userId, err := bg.validateUser(r)
	if err != nil {
		Error(w, http.StatusUnauthorized, err, err.Error())
		return
	}

	// parse the blog from body
	var blogPayload blogRequestPayload
	var blog entity.Blog
	err = json.NewDecoder(r.Body).Decode(&blogPayload)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "bad request body")
		return
	}

	// get the blog
	blog = blogPayload.Blog
	blog.UserId = userId
	// get the tags
	tags := blogPayload.Tags

	// TODO: process the blog before saving into the DB

	savedBlog, err := bg.blogApp.Save(blog, userId)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "unable to save the blog")
		return
	}

	// create the tag if its a new tag
	for _, t := range tags {
		ta, err := bg.tagApp.GetTagByName(t.Name)
		if err != nil {
			logrus.Warn("internal server error : " + err.Error())
		}
		if ta == nil { // tag do not exists create a new tag

			savedTag, err := bg.tagApp.Save(t)
			if err != nil {
				// TODO: create a retry background task OR push it into an event bus
				logrus.Warn(fmt.Sprintf("Tag %s is not saved", t.Name))
			}

			// save the blog-tag mapping
			var blogTagMapping entity.BlogTag
			blogTagMapping.BlogId = savedBlog.ID
			blogTagMapping.TagId = savedTag.ID
			err = bg.blogTagApp.Save(blogTagMapping)
			if err != nil {
				// TODO: Create retry event or PUsh it into an event bus
				logrus.Warn(fmt.Sprintf("Blog-Tag relation is not saved"))
			}
		}
	}

	JSON(w, http.StatusOK, "")
}

// user auth required
func (bg *Blog) Delete(w http.ResponseWriter, r *http.Request) {
	// validate user
	userId, err := bg.validateUser(r)
	if err != nil {
		Error(w, http.StatusUnauthorized, err, err.Error())
		return
	}
	// parse the blog from body
	var blog entity.Blog
	err = json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "bad request body")
		return
	}

	// check if the blog belongs to the current user
	bl, err := bg.blogApp.GetBlogById(blog.ID)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "couldn't find the blog")
		return
	}
	// check if the blog is of the same user
	if bl.UserId != userId {
		Error(w, http.StatusUnauthorized, fmt.Errorf("not your blog"), "blog is not yours")
		return
	}
	// delete the blog
	err = bg.blogApp.Delete(blog.ID)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "couldn't delete")
		return
	}
	// TODO: create event to delete the mapping of blog-tag
	// temprary : delete mapping from blog-tag table
	err = bg.blogTagApp.DeleteByBlogId([]uint64{blog.ID})
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "don't worry your blog is deleted")
		// TODO: we don't want to throw and error it is not vary important
	}
	JSON(w, http.StatusOK, "deleted")
}

func (bg *Blog) GetBlogById(w http.ResponseWriter, r *http.Request) {
	// validate user
	_, err := bg.validateUser(r)
	if err != nil {
		Error(w, http.StatusUnauthorized, err, err.Error())
		return
	}

	urlVars := mux.Vars(r)
	id, ok := urlVars["id"]
	idInt, err := strconv.ParseUint(id, 10, 64)
	if !ok || err != nil {
		Error(w, http.StatusBadRequest, fmt.Errorf("id not valid"), "blogid is not valid")
		return
	}

	blog, err := bg.blogApp.GetBlogById(idInt)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "blog not found")
		return
	}

	// TODO: this is service requirement ? what should we do ?
	// get all the tag ids for blog
	tags, err := bg.tagService.GetTagsByBlogId(blog.ID)
	// TODO: Do not throw error if tag service is down
	if err != nil {
		logrus.Warn(err.Error())
	}

	blogResp := blogResponse{
		Blog: *blog,
		Tags: tags,
	}

	JSON(w, http.StatusOK, blogResp)
}

func (bg *Blog) GetBlogByUserId(w http.ResponseWriter, r *http.Request) {

}

func (bg *Blog) GetBlogsByTagName(w http.ResponseWriter, r *http.Request) {
	// get limit
	limitQ := r.URL.Query().Get("limit")
	var limitInt int
	var err error
	if limitQ == "" {
		limitInt = 20 // default limit
	} else {
		limitInt, err = strconv.Atoi(limitQ)
		if err != nil {
			Error(w, http.StatusBadRequest, err, "limit arg is not correct")
		}
	}

	// get offset
	offsetQ := r.URL.Query().Get("offset")
	var offsetInt int
	if offsetQ != "" {
		offsetInt, err = strconv.Atoi(offsetQ)
		if err != nil {
			Error(w, http.StatusBadRequest, err, "offset arg is not correct")
		}
	}

	// get tagname
	tagName, ok := mux.Vars(r)["tag"]
	// TODO: validate tagname
	if !ok {
		Error(w, http.StatusBadRequest, fmt.Errorf("incorrect tagname"), "wrong url parameters")
		return
	}

	// get tag id from tagApp
	tag, err := bg.tagApp.GetTagByName(tagName)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "tag not found")
		return
	}
	// get blog-tag mapping from blogTagApp
	blogTagMapping, err := bg.blogTagApp.GetByTagId(tag.ID)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "internal error while fetching the blogs")
		return
	}

	blogIds := make([]uint64, 0)
	for _, i := range blogTagMapping {
		blogIds = append(blogIds, i.BlogId)
	}

	// get blogs using blog ids from blogApp
	blogs, err := bg.blogApp.GetBlogsByIds(blogIds, limitInt, offsetInt)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "blog not found for given tag")
		return
	}

	// blog response
	blogResp := make([]blogResponse, 0)
	for _, i := range blogs {
		tags, err := bg.tagService.GetTagsByBlogId(i.ID)

		if err != nil {
			// TODO: do not throw error if tag service is down
			// make array empty
			tags = []entity.Tag{}
		}

		blogResp = append(blogResp, blogResponse{
			Blog: i,
			Tags: tags,
		})
	}

	JSON(w, http.StatusOK, blogResp)
}

func (bg *Blog) GetBlogs(w http.ResponseWriter, r *http.Request) {
	// get limit
	limitQ := r.URL.Query().Get("limit")
	var limitInt int
	var err error
	if limitQ == "" {
		limitInt = 20 // default limit
	} else {
		limitInt, err = strconv.Atoi(limitQ)
		if err != nil {
			Error(w, http.StatusBadRequest, err, "limit arg is not correct")
		}
	}

	// get offset
	offsetQ := r.URL.Query().Get("offset")
	var offsetInt int
	if offsetQ != "" {
		offsetInt, err = strconv.Atoi(offsetQ)
		if err != nil {
			Error(w, http.StatusBadRequest, err, "offset arg is not correct")
		}
	}

	var blogs []entity.Blog
	blogs, err = bg.blogApp.GetBlogs(limitInt, offsetInt)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	blogResp := make([]blogResponse, 0)
	for _, i := range blogs {
		tags, err := bg.tagService.GetTagsByBlogId(i.ID)

		if err != nil {
			// TODO: do not throw error if tag service is down
			// make array empty
			tags = []entity.Tag{}
		}

		blogResp = append(blogResp, blogResponse{
			Blog: i,
			Tags: tags,
		})
	}

	JSON(w, http.StatusOK, blogResp)
}

// internal method
func (bg *Blog) validateUser(r *http.Request) (userId uuid.UUID, err error) {
	// get user's metadata
	userMeta, err := bg.tk.ExtractTokenMetadata(r)
	if err != nil {
		return
	}
	// check token
	userId, err = bg.au.FetchAuth(userMeta.TokenUuid)
	if err != nil {
		return
	}
	// check user into db
	_, err = bg.userApp.GetById(userId)
	if err != nil {
		return
	}
	return userId, err
}
