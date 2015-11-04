/* global $, wisply */
/**
 * @file Encapsulates the functionality for institution home page
 * @author Cristian Sima
 */
/**
 * @namespace AdminInstitutionsInstitutionHomeModule
 */
var AdminInstitutionsInstitutionHomeModule = function() {
	'use strict';

	/**
	 * The constructor does nothing
	 * @class Manager
	 * @memberof AdminInstitutionsInstitutionHomeModule
	 * @classdesc It manages the operations for the program
	 * @param {object} institution The object which contains the information about institution
	 * the current program
	 */
	var Manager = function Manager(institution) {
		this.institution = institution;
	};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends AdminInstitutionsInstitutionHomeModule.Manager */
		{
			/**
			 * It activates the listeners for all delete buttons
			 * @fires AccountsManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".deleteProgramButton").click(function(event) {
					event.preventDefault();
					var object,
						id;
					object = $(this);
					id = object.data("id");
					instance.confirmDelete(id, "program");
				});
				$(".deleteModuleButton").click(function(event) {
					event.preventDefault();
					var object,
						id;
					object = $(this);
					id = object.data("id");
					instance.confirmDelete(id, "module");
				});
			},
			/**
			 * It shows the confirmation dialog
			 * @param  {number} id  The ID of the program
			 * @param  {string} type  The type of the item ("program", "module")
			 */
			confirmDelete: function(id, type) {
				var msg = this.getDialogMessage(id, type);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguments for the confimation dialog
			 * @param  {number} id The ID of the program
			 * @param  {string} type  The type of the item ("program", "module")
			 * @return {object}         The arguments of the confimation message
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
					message: "The <strong>" + type + "</strong> will be removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets a definition
			 * @param  {number} id The ID of the program
			 * @param  {string} type  The type of the item ("program", "module")
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
					"url": '/admin/institutions/' + instance.institution.id + "/" + type + "/" + id + "/delete",
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
	var module = new AdminInstitutionsInstitutionHomeModule();
	wisply.loadModule("admin-institutions-institution-home", module);
});
