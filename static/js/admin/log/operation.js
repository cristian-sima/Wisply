/* global jQuery,$, wisply, bootbox */
/**
 * @file Encapsulates the functionality for managing the operations
 * @author Cristian Sima
 */
/**
 * @namespace Operations
 */
var Operations = function() {
	'use strict';

	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof Operations
	 * @classdesc It encapsulets the functions for operations
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends Operations.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 */
			init: function() {
				$(".see-full-explication").click(showExplication);
			}
		};
	/**
	 * It opens a bootstrap dialog and shows the explication
	 * @param [event] event The event which is generated
	 */
	function showExplication(event) {
		event.preventDefault();
		var instance,
			explication;
		instance = $(this);
		explication = instance.data("explication");
		wisply.message.show("Task Explication", explication);
	}
	return {
		Manager: Manager,
	};
};
jQuery(document).ready(function() {
	"use strict";
	var module = new Operations();
	wisply.operationsManager = new module.Manager();
	wisply.operationsManager.init();
});