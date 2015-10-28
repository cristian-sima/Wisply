/* global jQuery, document, $, wisply*/
var RegisterModule;

/**
* @file Encapsulates the functionality for RegisterModule page
* @author Cristian Sima
*/

/**
* @namespace RegisterModule
*/
var RegisterModule = function () {
  'use strict';

  /**
  * It does nothing important
  * @class Form
  * @memberof RegisterModule
  * @classdesc It encapsulates the functionality for RegisterModuleing an account
  */
  var Form = function Form() {
  };
  Form.prototype =
  /**
  * @lends RegisterModule.Form
  */
  {
    /**
    * It activates the listeners and focuses the name
    */
    init : function () {
      var info = wisply.getModule("server").getData();
      this.loadListeners();
      this.checkConfirmPassword();
      this.focusName();
      wisply.preloadLoadingImage();
      if(info.hasCaptcha) {
        var module = wisply.getModule("captcha"),
          name = "register-form-page";
        new module.Captcha({
          ID: info.ID,
          name: name,
        }).show();
      }
    },
    /**
    * It activates the listeners for form submitted and password focused
    */
    loadListeners: function() {
      this.passwordCompletedListener();
      this.submittedListener();
    },
    /**
    * It is called when the user clicks on the back button of the browser (or in the case Wisply detects problems with the form and the user goes back to the form). In case the value of the password is not empty, it shows the confirmation password field
    */
    checkConfirmPassword : function () {
      if($("#RegisterModule-password-confirm").val() !== "") {
        $("#div-confirm-password").show();
      }
    },
    /**
    * It focuses the name field
    */
    focusName: function () {
      $("#RegisterModule-name").focus();
    },
    /**
    * It is called when the form has been submitted. It shows the loading button
    */
    submittedListener : function() {
      $("#RegisterModule-form").on("submit", this.FireSubmited);
    },
    /**
    * It is called when the RegisterModule form has been submitted. It checks if the confirmation password is the same as the password. If so, it submits the form, else it shows a message
    * @param  {Event} event The event which is generated
    */
    FireSubmited: function(event) {
      event.preventDefault();
      var password = $('#RegisterModule-password').val(),
      confirmationPassword = $("#RegisterModule-password-confirm").val();
      if (password === confirmationPassword) {
        RegisterModule.showLoading();
        this.submit();
      } else {
        RegisterModule.showPasswordsDoNotMatch();
      }
    },
    /**
    * It shows the confirmation password field
    */
    passwordCompletedListener : function () {
      $("#RegisterModule-password").focus(function() {
        $("#div-confirm-password").show();
      });
    },
    /**
    * It shows the loading image
    */
    showLoading: function() {
      wisply.showLoading('#RegisterModule-submit-div', "medium");
    },
    /**
    * It tells the account that the passwords do not match
    */
    showPasswordsDoNotMatch: function () {
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
    return {
      Form: Form
    };
  };
  $(document).ready(function() {
    "use strict";
    var module = wisply.loadModule("register", new RegisterModule()),
      register = new module.Form();
    register.init();
  });
