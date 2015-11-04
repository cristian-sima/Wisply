/* global $, wisply */
/**
 * @file Encapsulates the functionality for managing the advance options for
 * subjects
 * @author Cristian Sima
 */
/**
 * @namespace AdminAdanceOptionsSubjectModule
 */
var AdminAdanceOptionsSubjectModule = function() {
	'use strict';
	/**
	 * The constructor saves the id and the name
	 * @class Subject
	 * @memberof AdminAdanceOptionsSubjectModule
	 * @classdesc It represents a subject
	 * @param {object} info It contains the information regarding the subject (id, name)
	 */
	var Subject = function Subject(info) {
		this.id = info.id;
		this.name = info.name;
	};
	/**
	 * The constructor activates the listeners
	 * @memberof AdminAdanceOptionsSubjectModule
	 * @class Manager
	 * @classdesc It encapsulets the functionality for the subjects
	 */
	var Manager = function Manager() {};
	Manager.prototype =
		/** @lends AdminAdanceOptionsSubjectModule.Manager */
		{
			/**
			 * It activates the listeners
			 */
			init: function() {
				var instance = this;
				$(".deleteSubjectButton").click(function(event) {
					event.preventDefault();
					var object, subject;
					object = $(this);
					subject = new Subject({
						id: object.data("id"),
						name: object.data("name"),
					});
					instance.confirmDelete(subject);
				});
			},
			/**
			 * It is called when the user wants to delete a subject. It asks for confirmation
			 * @param  {Subject} subject The reference to the subject object
			 */
			confirmDelete: function(subject) {
				/**
				 * It focuses the password input after the promt is shown
				 */
				function focusPassword() {
					setTimeout(function() {
						$("#promt-password").focus();
					}, 500);
				}
				var msg = this.getDeleteDialogMessage(subject);
				wisply.message.dialog(msg);
				focusPassword();
			},
			/**
			 * It returns the object which contain the arguments for the confirmation dialog
			 * @param  {string} type The type of action: "delete", "clear"
			 * @return {Object}        The arguements for dialog
			 * @see http://bootboxjs.com/
			 */
			getDeleteDialogMessage: function(subject) {
				var buttons,
					cancelButton,
					msg,
					mainButton,
					instance = this,
					subjectCopy = subject;
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
						var password, title;
						title = "Removing the subject <strong>" + subjectCopy.name + "</strong>";
						password = $("#promt-password").val();
						wisply.message.tellToWait(title);
						instance.delete(subjectCopy, password);
					},
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "Type your password",
					buttons: buttons,
					onEscape: true,
					message: '<input class="bootbox-input bootbox-input-text form-control" autocomplete="off" type="password" id="promt-password" />',
				};
				return msg;
			},
			/**
			 * It delets a subject
			 * @param  {Subject} subject The subject object
			 * @param {string} password The password from the user
			 */
			delete: function(subject, password) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * The callback called when the subject has been deleted. It shows a message and reloads the page in 2 seconds
				 */
				successCallback = function() {
					//wisply.message.showSuccess("The subject has been removed! Refreshing page...");
					window.location = "/admin/education";
				};
				/**
				 * The callback called when there was a problem. It shows a message
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					url: '/admin/education/subjects/' + subject.id + "/delete",
					success: successCallback,
					error: errorCallback,
					data: {
						password: password,
					},
				};
				wisply.executePostAjax(request);
			},
		};
	return {
		Subject: Subject,
		Manager: Manager,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AdminAdanceOptionsSubjectModule();
	wisply.loadModule("admin-education-subject-advance-options", module);
});
