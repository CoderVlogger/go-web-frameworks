# go-web-frameworks

Example web application projects with top popular and recently active Go lang web frameworks and toolkits.

I will implement the same features in all frameworks and compare them occasionally.

Currently checking:

- [gin](https://github.com/gin-gonic/gin)
- [echo](https://github.com/labstack/echo)
- [iris](https://github.com/kataras/iris)
- [fiber](https://github.com/gofiber/fiber)
- [micro](https://github.com/micro/micro)
- [goa](https://github.com/goadesign/goa)
- [gorilla](https://github.com/gorilla/);

## Related Resources

### Live Streams

I host live streams on [YouTube][codervlogger-youtube-live], [Twitch][codervlogger-twitch-live], and [Facebook Watch][codervlogger-facebook-live] about building this project. Follow me on [Twitter][codervlogger-twitter] for announcements.

### Accompanied Blog Post

Three is also a dedicated blog post where I gather all related tutorials and learning materials in one place. Learn more: [Go Web Frameworks and Toolkits][codervlogger-goweb]

## Features

List of checked features:

- "Hello, World!" and simple JSON response;
- CRUD JSON API for [the example entity][folder-pkg];
- Paging and error handling;
- Parsing query params;
- Accepting and parsing form data;
- Exposing static assets via an endpoint;
- Building website pages with templates;
- MVC Pattern (model, view, controller);
- Authentication and Authorization;

Additionaly, some examples with:
- Logging, tracing, and metrics;
- Docker and Docker Compos;
- Database and ORM;
- Work with MySQL;
- Work with PostgreSQL;
- Work with MongoDB;
- Work with Redis;
- Work with Kafka;

## Current Progress

### gin

Not started yet.

### echo

| Status | Feature |
|-|-|
| 👨🏼‍💻 | "Hello, World!" and simple JSON response |
| 👨🏼‍💻 | CRUD JSON API for [the example entity][folder-pkg] |
| 👨🏼‍💻 | Paging and error handling |
| 👨🏼‍💻 | Parsing query params |
| 👨🏼‍💻 | Accepting and parsing form data |
| 👨🏼‍💻 | Exposing static assets via an endpoint |
| 👨🏼‍💻 | Building website pages with templates |
| 👨🏼‍💻 | MVC Pattern (model, view, controller) |
| 👨🏼‍💻 | Authentication and Authorization |

### iris

| Feature                       | Status | ・ | Feature                       | Status |
|-------------------------------|--------|---|-------------------------------|--------|
| CRUD Endpoints                | [ ✅ ] | ・ | Log / Trace / Metrics         | [ 👨🏼‍💻 ] |
| Form / File Upload            | [ 👨🏼‍💻 ] | ・ | Docker / Docker Compose       | [ 👷🏼‍♂️ ] |
| Auth / JWT                    | [ 👨🏼‍💻 ] | ・ | Database / ORM                | [ 👨🏼‍💻 ] |
| Security / CORS               | [ 👨🏼‍💻 ] | ・ | Work with MySQL               | [ 👨🏼‍💻 ] |
| Optimization / GZIP           | [ 👨🏼‍💻 ] | ・ | Work with PostgreSQL          | [ 👨🏼‍💻 ] |
| Testing                       | [ 👨🏼‍💻 ] | ・ | Work with MongoDB             | [ 👨🏼‍💻 ] |
| Load testing / Benchmark      | [ 👨🏼‍💻 ] | ・ | Work with Redis               | [ 👨🏼‍💻 ] |
| Documentation / Swagger       | [ 👨🏼‍💻 ] | ・ | Work with Kafka               | [ 👨🏼‍💻 ] |

### fiber

Not started yet.

### micro

Not started yet.

### goa

Not started yet.

### gorilla

Not started yet.

[folder-pkg]: https://github.com/CoderVlogger/go-web-frameworks/tree/main/pkg "Shared package"
[codervlogger-goweb]: https://www.codervlogger.com/go-web-frameworks-and-toolkits/ "Go Web Frameworks and Toolkits"
[codervlogger-youtube-live]: https://www.youtube.com/channel/UCYsloGIkGmSzk2pw6yf6_Gw/live "YouTube"
[codervlogger-twitch-live]: https://www.twitch.tv/codervlogger "Twitch"
[codervlogger-facebook-live]: https://www.facebook.com/CoderVlogger/live_videos/ "Facebook Watch"
[codervlogger-twitter]: https://twitter.com/KenanBekk "Twitter"
