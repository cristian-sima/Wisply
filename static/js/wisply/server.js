/* global $, wisply */
/**
 * @file Encapsulates the functionality server module
 * @author Cristian Sima
 */
/**
 * @namespace ServerModule
 * The server module keep the data from server
 */
var ServerModule = function(data) {
	'use strict';
	/**
	 * The constructor creates a buffer using the parameter
	 * @class Manager
	 * @memberof ServerModule
	 * @classdesc The Manager for server data
	 * @param {object} o The default data from server
	 */
	function Manager(o) {
		this._data = o;
	}
	Manager.prototype =
		/** @lends Manager.ServerModule */
		{
			/**
			 * It adds data to the object. It removes the previous data
			 * @param {object} o The new object
			 */
			set: function(o) {
				this._data = o;
			},
			/**
			 * It returns the data that it contains
			 * @return {object} The data
			 */
			getData: function() {
				return this._data;
			}
		};
	return new Manager(data);
};
$(document).ready(function() {
	"use strict";
	wisply.loadModule("server", new ServerModule());
});
