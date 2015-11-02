package captcha

import (
	"fmt"
	"strconv"
	"time"
)

type action struct {
	page      string
	timestamp int64
	ip        string
	count     int
}

func (action action) isExpired() bool {
	for _, currentAction := range currentListOfActions.actions {
		fmt.Println(currentAction.page + " " + strconv.Itoa(currentAction.count))
	}
	allowedTime := getAllowedTime(action.page)
	return action.timestamp+allowedTime < getCurrentTimestamp()
}

func (action action) requireCaptcha() bool {
	fmt.Println("allowed times: " + strconv.Itoa(getAllowedNumber(action.page)))
	return action.count > getAllowedNumber(action.page)
}

// List is the list of actions on the server
type List struct {
	actions []*action
}

func (list *List) addAction(page, ip string) {
	var (
		exists = false
	)
	for index, action := range list.actions {
		if action.ip == ip && action.page == page {
			action.count++
			exists = true
		}
		if action.isExpired() {
			list.actions = append(list.actions[:index], list.actions[index+1:]...)
		}
	}
	if !exists {
		item := &action{
			page:      page,
			ip:        ip,
			count:     1,
			timestamp: getCurrentTimestamp(),
		}
		list.actions = append(list.actions, item)
	}
}

func newList() {
	items := []*action{}
	currentListOfActions = List{
		actions: items,
	}
}

// RequireCaptcha checks if the combination of page and ip needs a captcha
//
func RequireCaptcha(page, ip string) bool {
	for _, action := range currentListOfActions.actions {
		if action.ip == ip && action.page == page {
			return action.requireCaptcha()
		}
	}
	return false
}

// RegisterAction adds a new action for a page
func RegisterAction(page, ip string) {
	currentListOfActions.addAction(page, ip)
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
