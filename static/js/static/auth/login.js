/* global $, wisply */
var login;

/**
* @file Encapsulates the functionality for login page
* @author Cristian Sima
*/

/**
 * Requires CaptchaModule, ServerModule
* @namespace LoginModule
*/
var LoginModule = function () {
  'use strict';

  /**
  * Encapsulates the functionality for login
  * @class Form
  * @memberof LoginModule
  * @classdesc It represents the login gateway
  */
  var Form = function Form() {
  };
  Form.prototype =
  /** @lends LoginModule.Form */
  {
    /**
    * It loads the listeners and focuses the name
    */
    init : function () {
      var info = wisply.getModule("server").getData();
      this.loadListeners();
      this.focusName();
      wisply.preloadLoadingImage();
      if(info.hasCaptcha) {
        var module = wisply.getModule("captcha"),
          name = "login-form-captcha";
        new module.Captcha({
          element: $("#" + name),
          ID: info.ID,
          name: name,
        }).show();
      }
    },
    /**
    * It adds a listener for form submit
    */
    loadListeners: function() {
      $("#login-form").on("submit", this.FireSubmited);
    },
    /**
    * It focuses the name field.
    */
    focusName: function () {
      $("#login-email").focus();
    },
    /**
    * It is fired when the form has been submitted. It shows the loading image
    */
    FireSubmited: function() {
      wisply.showLoading('#login-submit-div', "medium");
    }
  };
  return {
      Form: Form
  };
};
$(document).ready(function() {
  "use strict";
  var module = wisply.loadModule("login", new LoginModule());
  login = new module.Form();
  login.init();
});
