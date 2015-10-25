/* global jQuery,$, wisply */
/**
 * @file Encapsulates the functionality for managing the accounts
 * @author Cristian Sima
 */
/**
 * @namespace Accounts
 */
var Accounts = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class
	 * @memberof Accounts
	 * @classdesc It represents a Wisply account
	 * @param {number} id   The id of the account
	 * @param {string} name The name of the account
	 */
	var Account = function Account(id, name) {
		this.id = id;
		this.name = name;
	};
	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof Accounts
	 * @classdesc It encapsulets the functions for accounts.
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends Accounts.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 * @fires AccountsManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".deleteAccountButton").click(function(event){
					event.preventDefault();
					var object,
						name,
						id,
						account;
					object = $(this);
					id = object.data("id");
					name = object.data("name");
					account = new Account(id, name);
					instance.confirmDelete(account);
				});
			},
			/**
			 * It requests the user to confirm
			 * @param  {Account} account  The Account object to be deleted
			 */
			confirmDelete: function(account) {
				var msg = this.getDialogMessage(account);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements of the confimation message
			 * @param  {Account} account The Account object to be deleted
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(account) {
				var buttons,
					cancelButton,
					msg,
					mainButton;
				cancelButton = {
					label: "Cancel",
					className: "btn-success",
					callback: function() {
						this.modal('hide');
					}
				};
				mainButton = {
					label: "Delete account",
					className: "btn-danger",
					callback: function() {
						wisply.accountsManager.delete(account);
					}
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "We need your confirmation",
					message: "The account <b>" + account.name + "</b> will be permanently removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets a user
			 * @param  {Account} account The account to be deleted
			 */
			delete: function(account) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when the deletion has been done. It reloads the page in 2 seconds
				 * @ignore
				 */
				successCallback = function() {
					wisply.message.showSuccess("The account has been removed! Refreshing page...");
					wisply.reloadPage();
				};
				/**
				 * It is called when there has been problems
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					"url": '/admin/accounts/delete/' + account.id,
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
			}
		};
	return {
		Manager: Manager,
		Account: Account
	};
};
jQuery(document).ready(function() {
	"use strict";
	var module = new Accounts();
	wisply.accountsManager = new module.Manager();
	wisply.accountsManager.init();
});
