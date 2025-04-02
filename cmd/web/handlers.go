package main

import (
	"app/pkg/logic"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var GlobalLoginUser string

func SessionCheck(c *gin.Context) {
	s := sessions.Default(c)
	if user := s.Get("User"); user != nil {
		if GlobalLoginUser == "" {
			GetExit(c)
			location := url.URL{Path: "/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		res, err := logic.GetUser(GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["User"] = res
	}
}

func GetResult(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		user, err := logic.GetToId_sql(GlobalLoginUser)
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "error.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		event_ := user.Events[c.Param("Name")]
		oneTiket, summAll, summPrc, moneyUp, err := logic.EventCalculation(event_.VipPrice, event_.StandartPrice, event_.EconomPrice, event_.CountTicketVip, event_.CountTicketStandart, event_.CountTicketEconom, event_.NetProfit, event_.Expenses, "noSave", "", "", GlobalLoginUser)
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		c.Keys["oneTiket"] = oneTiket
		c.Keys["summAll"] = summAll
		c.Keys["summPrc"] = summPrc
		c.Keys["moneyUp"] = moneyUp
		c.HTML(http.StatusOK, "result.html", c.Keys)
		c.Keys["oneTiket"] = nil
		c.Keys["summAll"] = nil
		c.Keys["summPrc"] = nil
		c.Keys["moneyUp"] = nil
		return
	}
	location := url.URL{Path: "/lk"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func PostHome(c *gin.Context) {
	session := sessions.Default(c)
	if c.Request.Method == http.MethodPost {
		countTicketVip, err := strconv.Atoi(c.PostForm("count-ticket-vip"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		countTicketStandart, err := strconv.Atoi(c.PostForm("count-ticket-standart"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		countTicketEconom, err := strconv.Atoi(c.PostForm("count-ticket-econom"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		plusMoney, err := strconv.Atoi(c.PostForm("plus-money"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		minusMoney, err := strconv.Atoi(c.PostForm("minus-money"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}

		vipPrice, err := strconv.Atoi(c.PostForm("vip-price"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		standartPrice, err := strconv.Atoi(c.PostForm("standart-price"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		economPrice, err := strconv.Atoi(c.PostForm("econom-price"))
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		checkSave := c.PostForm("check-save")
		if checkSave == "save" {
			if session.Get("User") == nil {
				c.Keys["error"] = "Войдите в аккаунт для сохранения"
				c.HTML(403, "index.html", c.Keys)
				c.Keys["error"] = nil
				return
			}
			if c.PostForm("name") == "" || c.PostForm("description") == "" {
				c.Keys["error"] = "Заполните все поля для сохранения"
				c.HTML(403, "index.html", c.Keys)
				c.Keys["error"] = nil
				return
			}
		}
		oneTiket, summAll, summPrc, moneyUp, err := logic.EventCalculation(vipPrice, standartPrice, economPrice, countTicketVip, countTicketStandart, countTicketEconom, plusMoney, minusMoney, checkSave, c.PostForm("name"), c.PostForm("description"), GlobalLoginUser)
		if err != nil {
			c.Keys["error"] = err
			c.HTML(403, "index.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
		c.Keys["oneTiket"] = oneTiket
		c.Keys["summAll"] = summAll
		c.Keys["summPrc"] = summPrc
		c.Keys["moneyUp"] = moneyUp
		c.HTML(http.StatusOK, "index.html", c.Keys)
		c.Keys["oneTiket"] = nil
		c.Keys["summAll"] = nil
		c.Keys["summPrc"] = nil
		c.Keys["moneyUp"] = nil
	}
}

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", c.Keys)
}

// func GetCourses(c *gin.Context) {
// 	res, err := logic.GetAllCourse_sql()
// 	if err != nil {
// 		c.Keys["error"] = err
// 		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
// 		c.Keys["error"] = nil
// 		return
// 	}
// 	c.Keys["res"] = res
// 	c.HTML(http.StatusOK, "courses.html", c.Keys)
// 	c.Keys["res"] = nil
// }

// func GetCoursesTarget(c *gin.Context) {
// 	session := sessions.Default(c)
// 	if session.Get("User") == nil {
// 		location := url.URL{Path: "/login"}
// 		c.Redirect(http.StatusFound, location.RequestURI())
// 	}
// 	num, _ := strconv.Atoi(c.Param("id"))
// 	data, err := logic.GetToCourseId_sql(num)
// 	if err != nil {
// 		c.Keys["error"] = err
// 		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
// 		c.Keys["error"] = nil
// 		return
// 	}
// 	c.Keys["data"] = data
// 	c.HTML(http.StatusOK, "course-target.html", c.Keys)
// 	c.Keys["data"] = nil
// }

// func PostCoursesTarget(c *gin.Context) {
// 	id_course := c.Param("id")
// 	if c.Request.Method == http.MethodPost {
// 		new_id_course, _ := strconv.Atoi(id_course)
// 		err := logic.AddFavoritesCourse(new_id_course, GlobalLoginUser)
// 		if err != nil {
// 			c.Keys["error"] = err
// 			c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
// 			c.Keys["error"] = nil
// 			return
// 		}
// 	}
// 	location := url.URL{Path: "/courses/" + id_course}
// 	c.Redirect(http.StatusFound, location.RequestURI())
// }

// Функция удаления
func PostLk(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		c.Keys["error"] = "Session nil"
		c.HTML(http.StatusOK, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	err := logic.DeleteFavorites_sql(c.PostForm("name"), GlobalLoginUser)
	if err != nil {
		c.Keys["error"] = err
		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	location := url.URL{Path: "/lk"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetLk(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		c.HTML(http.StatusOK, "lk.html", c.Keys)
		return
	}
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetExit(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		session.Set("User", nil)
		session.Save()
		GlobalLoginUser = ""
	}
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		location := url.URL{Path: "/lk"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	c.HTML(http.StatusOK, "login.html", c.Keys)
}

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	if c.Request.Method == http.MethodPost {
		user, err := logic.GetUser(c.PostForm("login"))
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{"error": err})
			return
		}
		if c.PostForm("password") != user.Password {
			c.HTML(http.StatusOK, "login.html", gin.H{"error": "Неправильный пароль"})
			return
		}
		if session.Get("User") == nil {
			session.Set("User", user)
			GlobalLoginUser = user.Login
			session.Save()
		}
		location := url.URL{Path: "/lk"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

// c.HTML(http.StatusOK, "index.html", gin.H{"id": c.Param("id")})

// var (
// 	name = c.PostForm("name")
// 	typs = c.PostForm("type")
// 	err  error
// )
// _, err = logic.ProdCreate(typs, name)
// c.HTML(http.StatusOK, "index.html", gin.H{"err": err})

// q := url.Values{}
// q.Set("id", c.Param("id"))
// location := url.URL{Path: "/dish/" + c.Param("id"), RawQuery: q.Encode()}
// c.Redirect(http.StatusFound, location.RequestURI())

// q := url.Values{}
// location := url.URL{Path: "/drink/" + c.Param("id"), RawQuery: q.Encode()}
// c.Redirect(http.StatusFound, location.RequestURI())
