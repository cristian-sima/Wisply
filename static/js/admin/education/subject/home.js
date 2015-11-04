/* global $, wisply */
/**
 * @file Encapsulates the functionality for process home page
 * @author Cristian Sima
 */
/**
 * @namespace AdminEducationSubjectHome
 */
var AdminEducationSubjectHome = function() {
	'use strict';

	/**
	 * The constructor does nothing
	 * @class Manager
	 * @memberof AdminEducationSubjectHome
	 * @classdesc It manages the operations for the subject
	 * @param {object} subject The object which contains the information about
	 * the current subject
	 */
	var Manager = function Manager(subject) {
		this.subject = subject;
	};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends AdminEducationSubjectHome.Manager */
		{
			/**
			 * It activates the listeners for all delete buttons
			 * @fires AccountsManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".deleteDefinitionButton").click(function(event) {
					event.preventDefault();
					var object,
						id;
					object = $(this);
					id = object.data("id");
					instance.confirmDelete(id, "definition");
				});
				$(".deleteKAButton").click(function(event) {
					event.preventDefault();
					var object,
						id;
					object = $(this);
					id = object.data("id");
					instance.confirmDelete(id, "ka");
				});
			},
			/**
			 * It shows the confirmation dialog
			 * @param  {number} id  The ID of the definition
			 * @param  {string} type  The type of item
			 */
			confirmDelete: function(id, type) {
				var msg = this.getDialogMessage(id, type);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements for the confimation dialog
			 * @param  {number} id The ID of the definition
			 * @param  {string} type  The type of item
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(id, type) {
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
						label: "Delete",
						className: "btn-danger",
						callback: function() {
							instance.delete(id, type);
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
					message: "The <strong>" + ((type === "ka")?"knowledge area":type) + "</strong> will be removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets a definition
			 * @param  {string} type  The type of item
			 * @param  {number} id The ID of the definition
			 */
			delete: function(id, type) {
				var request,
					successCallback,
					errorCallback,
					instance = this;
				/**
				 * It is called when the deletion has been done.
				 * @ignore
				 */
				successCallback = function() {
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
					"url": '/admin/education/subjects/' + instance.subject.id + "/"+ type + "/" + id + "/delete",
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
			}
		};
	return {
		Manager: Manager
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AdminEducationSubjectHome();
	wisply.loadModule("admin-education-subject-home", module);
});
