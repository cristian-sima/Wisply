/* global jQuery, wisply */
var login;

(function ($) {
  'use strict';
  function Login() {
    this.init();
  }
  Login.prototype = {
    init : function () {
      this.loadListeners();
    },
    loadListeners: function() {
      $("#login-form").on("submit", this.FireFormSubmited);
    },
    FireFormSubmited: function() {
      login.showLoading();
    },
    showLoading: function() {
      wisply.showLoading('#login-submit-div', "medium");
    }
  };
  function initLogin() {
    login = new Login();
  }
  $(document).ready(initLogin);
}(jQuery, wisply, login));
