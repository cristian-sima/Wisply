/* global jQuery,$, wisply */
/**
 * @file Encapsulates the functionality for managing the APITableList
 * @author Cristian Sima
 */
/**
 * @namespace APITableList
 */
var APITableList = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class
	 * @memberof APITableList
	 * @classdesc It represents a Wisply table
	 * @param {number} id   The id of the table
	 * @param {string} name The name of the table
	 */
	var Table = function Table(id, name) {
		this.id = id;
		this.name = name;
	};
	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof APITableList
	 * @classdesc It encapsulets the functions for APITableList.
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends APITableList.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 * @fires APITableListManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".download-table").click(function(event) {
					var name,	id,	table, object;
					event.preventDefault();
					object = $(this);
					id = object.data("id");
					name = object.data("name");
					table = new Table(id, name);
					instance.downloadTable(table);
				});
			},
			/**
			 * It tells the server to download the table and shows a waiting message
			 * @param  {APITableList.Table} table The table to be downloaded
			 */
			downloadTable: function(table) {
				var request,
					successCallback,
					errorCallback,
					box;
				/**
				 * It is called when the deletion has been done. It reloads the page in 2 seconds
				 * @ignore
				 */
				successCallback = function() {
					box.modal("hide");
					window.location.href = "/static/web/robots.txt";
				};
				/**
				 * It is called when there has been problems
				 * @ignore
				 */
				errorCallback = function() {
					box.modal("hide");
					wisply.message.showError("There was a problem with your request :(");
				};
				request = {
					"url": '/api/table/download/' + table.name,
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
				box = wisply.message.tellToWait("Downloading table <strong>" + table.name + "</strong> ...");
			},
		};
	return {
		Manager: Manager,
		Table: Table
	};
};
jQuery(document).ready(function() {
	"use strict";
	var module = new APITableList();
	wisply.APITableListManager = new module.Manager();
	wisply.APITableListManager.init();
});
