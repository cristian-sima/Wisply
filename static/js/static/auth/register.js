/* global jQuery, document, $, wisply*/
var register;

/**
* @file Encapsulates the functionality for register page
* @author Cristian Sima
*/

/**
* @namespace Register
*/
var Register = function () {
  'use strict';

  /**
  * It does nothing important
  * @class Form
  * @memberof Register
  * @classdesc It encapsulates the functionality for registering an account
  */
  var Form = function Form() {
  };
  Form.prototype =
  /**
  * @lends Register.Form
  */
  {
    /**
    * It activates the listeners and focuses the name
    */
    init : function () {
      this.loadListeners();
      this.checkConfirmPassword();
      this.focusName();
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
      if($("#register-password-confirm").val() !== "") {
        $("#div-confirm-password").show();
      }
    },
    /**
    * It focuses the name field
    */
    focusName: function () {
      $("#register-name").focus();
    },
    /**
    * It is called when the form has been submitted. It shows the loading button
    */
    submittedListener : function() {
      $("#register-form").on("submit", this.FireSubmited);
    },
    /**
    * It is called when the register form has been submitted. It checks if the confirmation password is the same as the password. If so, it submits the form, else it shows a message
    * @param  {Event} event The event which is generated
    */
    FireSubmited: function(event) {
      event.preventDefault();
      var password = $('#register-password').val(),
      confirmationPassword = $("#register-password-confirm").val();
      if (password === confirmationPassword) {
        register.showLoading();
        this.submit();
      } else {
        register.showPasswordsDoNotMatch();
      }
    },
    /**
    * It shows the confirmation password field
    */
    passwordCompletedListener : function () {
      $("#register-password").focus(function() {
        $("#div-confirm-password").show();
      });
    },
    /**
    * It shows the loading image
    */
    showLoading: function() {
      wisply.showLoading('#register-submit-div', "medium");
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
  jQuery(document).ready(function() {
    "use strict";
    var module = new Register();
    register = new module.Form();
    register.init();
  });
