/* global $, wisply */
/**
 * @file Encapsulates the functionality for managing the tables to download
 * @author Cristian Sima
 */
/**
 * @namespace AdminAPIModule
 */
var AdminAPIModule = function() {
	'use strict';
	/**
	 * The constructor saves the id and the name
	 * @class Table
	 * @memberof AdminAPIModule
	 * @classdesc It represents a table
	 * @param {object} info It contains the information regarding the table (id, name)
	 */
	var Table = function Table(info) {
		this.id = info.id;
		this.name = info.name;
	};
	/**
	 * The constructor does nothing
	 * @memberof AdminAPIModule
	 * @class List
	 * @classdesc It encapsulets the functionality for the tables
	 */
	var List = function List() {};
	List.prototype =
		/** @lends AdminAPIModule.List */
		{
			/**
			 * It activates the listeners
			 */
			init: function() {
				var instance = this;
				$(".makeTablePrivate").click(function(event) {
					event.preventDefault();
					var object,
						table;
					object = $(this);
					table = new Table({
						id: object.data("id"),
						name: object.data("name"),
					});
					instance.confirmDelete(table);
				});
				GUI.activateActionListeners();
			},
			/**
			 * It is called when the user wants to delete a table. It asks for confirmation
			 * @param  {Table} table The reference to the table object
			 */
			confirmDelete: function(table) {
				var msg = this.getDialogMessage(table);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the object which contain the arguments for the confirmation dialog
			 * @param  {Table} table The table object
			 * @return {Object}        The arguements for dialog
			 * @see http://bootboxjs.com/
			 */
			getDialogMessage: function(table) {
				var buttons,
					cancelButton,
					msg,
					mainButton,
					instance = this;
				cancelButton = {
					label: "No, thanks",
					className: "btn-success",
					callback: function() {
						this.modal('hide');
					}
				};
				mainButton = {
					label: "Yes",
					className: "btn-danger",
					callback: function() {
						instance.delete(table);
					}
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "We need your confirmation!",
					message: "Do you want to remove this table from the public list?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets a table
			 * @param  {Table} table The table object
			 */
			delete: function(table) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * The callback called when the table has been deleted. It shows a message and reloads the page in 2 seconds
				 */
				successCallback = function() {
					wisply.message.showSuccess("The table has been removed! Refreshing page...");
					wisply.reloadPage(2000);
				};
				/**
				 * The callback called when there was a problem. It shows a message
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					"url": '/admin/developer/delete',
					data: {
						"table-id": table.id,
					},
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
			}
		};
	/**
	 * The constructor does nothing
	 * @memberof AdminAPI
	 * @class GUI
	 * @classdesc It encapsulets the GUI functionality for download tables
	 */
	var GUI = function GUI() {};
	/**
	 * It activates all the listeners for the actions
	 */
	GUI.activateActionListeners = function() {
		$(".tables-init-harvest").click(function(event) {
			event.preventDefault();
			var object,
				id,
				form,
				xsrf;
			object = $(this);
			id = object.data("id");
			xsrf = wisply.getXSRF();
			form = '<form action="/admin/harvest/init/' + id + '" method="POST">' + '<input type="hidden" name="_xsrf" value="' + xsrf + '">' + '</form>';
			$(form).submit();
		});
		$(function() {
			$('[data-toggle="tooltip"]').tooltip();
		});
	};
	return {
		Table: Table,
		List: List,
		GUI: GUI
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AdminAPIModule();
	wisply.loadModule("admin-api", module);
});
