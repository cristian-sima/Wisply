/* global $, wisply */
/**
 * @file Encapsulates the functionality for managing the accounts
 * @author Cristian Sima
 */
/**
 * @namespace AdminAccounts
 */
var AdminAccountsModule = function() {
	'use strict';
	/**
	 * The constructor saves the name and the id of the account
	 * @class Account
	 * @memberof AdminAccounts
	 * @classdesc It represents a Wisply account
	 * @param {number} id   The id of the account
	 * @param {string} name The name of the account
	 */
	var Account = function Account(id, name) {
		this.id = id;
		this.name = name;
	};
	/**
	 * The constructor does nothing
	 * @class List
	 * @memberof AdminAccounts
	 * @classdesc It manages the operations for deleting the accounts
	 */
	var List = function List() {};
	/**
	 * @memberof Manager
	 */
	List.prototype =
		/** @lends AdminAccounts.List */
		{
			/**
			 * It activates the listeners for all delete buttons
			 * @fires AccountsManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".deleteAccountButton").click(function(event) {
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
			 * It shows the confirmation dialog
			 * @param  {Account} account  The Account object to be deleted
			 */
			confirmDelete: function(account) {
				var msg = this.getDialogMessage(account);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements for the confimation dialog
			 * @param  {Account} account The Account object to be deleted
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(account) {
				/**
				 * It returns the buttons for the dialog
				 * @return {object} The buttons for the dialog
				 */
				function getButtons() {
					var cancelButton,
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
							instance.delete(account);
						}
					};
					return {
						"cancel": cancelButton,
						"main": mainButton
					};
				}
				var msg,
					instance = this,
					buttons = getButtons();
				msg = {
					title: "We need your confirmation",
					message: "The account <b>" + account.name + "</b> will be permanently removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets an account
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
				 * It is called when the request has failed
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
		List: List,
		Account: Account
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AdminAccountsModule();
	wisply.loadModule("admin-accounts", module);
});
