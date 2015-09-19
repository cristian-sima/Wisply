/* global $, wisply */
var login;

/**
* @file Encapsulates the functionality for login page
* @author Cristian Sima
*/

/**
* @namespace Login
*/
var Login = function () {
  'use strict';

  /**
  * Encapsulates the functionality for login
  * @class Form
  * @memberof Login
  * @classdesc It represents the login gateway
  */
  var Form = function Form() {
  };
  Form.prototype =
  /** @lends Login.Form */
  {
    /**
    * It loads the listeners and focuses the name
    */
    init : function () {
      this.loadListeners();
      this.focusName();
      wisply.preloadLoadingImage();
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
  var module = new Login();
  login = new module.Form();
  login.init();
});
