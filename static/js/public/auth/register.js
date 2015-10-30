/* global document, $, wisply*/
var RegisterModule;
/**
 * @file Encapsulates the functionality for RegisterModule page
 * @author Cristian Sima
 */
/**
 * @namespace RegisterModule
 */
var RegisterModule = function() {
	'use strict';
	/**
	 * It does nothing important
	 * @class Form
	 * @memberof RegisterModule
	 * @classdesc It encapsulates the functionality for RegisterModuleing an account
	 */
	var Form = function Form() {};
	Form.prototype =
		/**
		 * @lends RegisterModule.Form
		 */
		{
			/**
			 * It activates the listeners and focuses the name
			 */
			init: function() {
				var info = wisply.getModule("server").getData();
				this.loadListeners();
				this.checkConfirmPassword();
				this.focusName();
				wisply.preloadLoadingImage();
				if (info.hasCaptcha) {
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
			checkConfirmPassword: function() {
				if ($("#register-password-confirm").val() !== "") {
					$("#div-confirm-password").show();
				}
			},
			/**
			 * It focuses the name field
			 */
			focusName: function() {
				$("#register-name").focus();
			},
			/**
			 * It is called when the form has been submitted. It shows the loading button
			 */
			submittedListener: function() {
				var instance = this;
				$("#register-form").on("submit", function(event){
					event.preventDefault();
					var password = $('#register-password').val(),
						confirmationPassword = $("#register-password-confirm").val();
					if (password === confirmationPassword) {
						RegisterModule.showLoading();
						this.submit();
					} else {
						instance.showPasswordsDoNotMatch();
					}
				});
			},
			/**
			 * It shows the confirmation password field
			 */
			passwordCompletedListener: function() {
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
			showPasswordsDoNotMatch: function() {
				var args = {
					title: "I got a problem",
					message: "The confirmation password is equal to the password. Correct this.",
					onEscape: true,
					buttons: {
						cancel: {
							label: "Ok",
							className: "btn-primary",
							callback: function() {
								this.modal('hide');
							}
						}
					}
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
	var module = new RegisterModule();
	wisply.loadModule("register", module);
});
