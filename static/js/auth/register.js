/* global jQuery, wisply, bootbox*/
var register;

(function ($) {
  'use strict';
  function Register() {
    this.init();
  }
  Register.prototype = {
    init : function () {
      this.loadListeners();
      this.checkConfirmPassword();
    },
    loadListeners: function() {
      this.passwordCompletedListener();
      this.formSubmittedListener();
    },
    checkConfirmPassword : function () {
      if($("#register-password-confirm").val() !== "") {
        $("#div-confirm-password").show();
      }
    },
    formSubmittedListener : function() {
      $("#register-form").on("submit", this.FireFormSubmited);
    },
    FireFormSubmited: function(event) {
      event.preventDefault();
      var password = $('#register-password').val(),
      confirmationPassword = $("#register-password-confirm").val();
      if (password === confirmationPassword) {
        register.showLoading();
        this.submit();
      } else {
        register.alertUser();
      }
    },
    passwordCompletedListener : function () {
      $("#register-password").focus(function() {
        $("#div-confirm-password").show();
      });
    },
    showLoading: function() {
      wisply.showLoading('#register-submit-div', "medium");
    },
    alertUser: function () {
      var args = {
        title: "I got a problem",
        message: "The confirmation password is equal to the password. Correct this.",
        onEscape: true,
        buttons: {
          cancel: {
            label: "Ok",
            className: "btn-default",
            callback: function() {
              this.modal('hide');
            }
          }}
        };
        wisply.message.alert(args);
    }
    };
    function initRegister() {
      register = new Register();
    }
    $(document).ready(initRegister);
  }(jQuery, wisply, bootbox, register));
