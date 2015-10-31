/* global jQuery,$, wisply */
/**
 * @file Encapsulates the functionality for managing the developerTableList
 * @author Cristian Sima
 */
/**
 * @namespace developerTableList
 */
var developerTableListModule = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class
	 * @memberof developerTableListModule
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
	 * @memberof developerTableListModule
	 * @classdesc It encapsulets the functions for developerTableListModule.
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends developerTableListModule.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 * @fires developerTableListManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".table-name").each(function(){
						var object = $(this);
						object.html(object.html().replace("_", " "));
				});
				$(".download-table").click(function(event) {
					var name,	id,	table, object;
					event.preventDefault();
					object = $(this);
					id = object.data("id");
					name = object.data("name");
					table = new Table(id, name);
					instance.downloadTable(table);
				});
				wisply.activateTooltip();
				$('body').popover({ selector: '[data-popover]', trigger: 'click hover', placement: 'auto', delay: {show: 50, hide: 400}});
			},
			/**
			 * It tells the server to download the table and shows a waiting message
			 * @param  {developerTableListModule.Table} table The table to be downloaded
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
					window.location.href = '/developer/table/download/' + table.name;
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
					"url": '/developer/table/generate/' + table.name,
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
				box = wisply.message.tellToWait("Downloading table <strong>" + table.name + "</strong> ...");
			},
		};

	/*
			Credits: http://jsfiddle.net/wojtekkruszewski/zf3m7/22/
	 */
	var originalLeave = $.fn.popover.Constructor.prototype.leave;
	$.fn.popover.Constructor.prototype.leave = function(obj){
	  var self = obj instanceof this.constructor ?
	    obj : $(obj.currentTarget)[this.type](this.getDelegateOptions()).data('bs.' + this.type);
	  var container, timeout;

	  originalLeave.call(this, obj);

	  if(obj.currentTarget) {
	    container = $(obj.currentTarget).siblings('.popover');
	    timeout = self.timeout;
	    container.one('mouseenter', function(){
	      //We entered the actual popover – call off the dogs
	      clearTimeout(timeout);
	      //Let's monitor popover content instead
	      container.one('mouseleave', function(){
	        $.fn.popover.Constructor.prototype.leave.call(self, self);
	      });
	    });
	  }
	};
	/*
		End credits
	 */
	return {
		Manager: Manager,
		Table: Table
	};
};
jQuery(document).ready(function() {
	"use strict";
	var module = new developerTableListModule();
	wisply.loadModule("developer-table-list", module);
});
