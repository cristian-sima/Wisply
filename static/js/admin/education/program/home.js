/* global $, wisply */
/**
 * @file Encapsulates the functionality for process home page
 * @author Cristian Sima
 */
/**
 * @namespace AdminEducationProgramHome
 */
var AdminEducationProgramHome = function() {
	'use strict';

	/**
	 * The constructor does nothing
	 * @class Manager
	 * @memberof AdminEducationProgramHome
	 * @classdesc It manages the operations for the program
	 * @param {object} program The object which contains the information about
	 * the current program
	 */
	var Manager = function Manager(program) {
		this.program = program;
	};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends AdminEducationProgramHome.Manager */
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
					instance.confirmDelete(id);
				});
			},
			/**
			 * It shows the confirmation dialog
			 * @param  {number} id  The ID of the definition
			 */
			confirmDelete: function(id) {
				var msg = this.getDialogMessage(id);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements for the confimation dialog
			 * @param  {number} id The ID of the definition
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(id) {
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
							instance.delete(id);
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
					message: "The definition will be removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets a definition
			 * @param  {number} id The ID of the definition
			 */
			delete: function(id) {
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
					"url": '/admin/education/programs/' + instance.program.id + "/definition/" + id + "/delete",
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
	var module = new AdminEducationProgramHome();
	wisply.loadModule("admin-education-program-home", module);
});
