FROM alpine:latest
COPY --from=build /app/app .
EXPOSE 8000:8000

CMD [ "./app" ]