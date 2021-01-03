module awesomeProject

go 1.13

require (
	awesomeProject/ES/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/chromedp/cdproto v0.0.0-20200116234248-4da64dd111ac
	github.com/chromedp/chromedp v0.5.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/jmoiron/sqlx v1.2.0
	github.com/juju/ratelimit v1.0.1
	github.com/labstack/gommon v0.3.0

	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	github.com/lxn/win v0.0.0-20191128105842-2da648fda5b4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.20.11+incompatible
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/zxmrlc/log v0.0.0-20200612082315-9e0c7ff11ddb
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/tools v0.0.0-20200426102838-f3a5411a4c3b // indirect
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.57.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	nhooyr.io/websocket v1.8.6

)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1 // indirect

replace awesomeProject/ES/kubernetes/staging/src/k8s.io/apimachinery => k8s.io/apimachinery v0.18.12-rc.1
