/* global $, Harvest, wisply, harvestHistory */
/**
 * @file Encapsulates the functionality for list harvesting
 * @author Cristian Sima
 */
/**
 * @namespace HarvestListModule
 * It holds the functionality to see a live list of repositories
 */
var HarvestListModule = function() {
	'use strict';
	var Stages = [{
		id: 0,
		name: "Get info from server",
		perform: function(manager) {
			this.init();
			this.manager = manager;
			this.manager.sendMessage("get-all-status");
			this.manager.firedStageFinished();
			harvestHistory.log("Listening for changes...");
		},
		init: function() {
			harvestHistory.setGUI("#harvest-history-element");
			$("#harvest-history-button").click(function() {
				$("#harvest-history-container").modal('show');
				harvestHistory.gui.activate();
			});
			$("#harvest-history-container").on('hidden.bs.modal', function() {
				harvestHistory.gui.disable();
			});
		}
	}];
	/**
	 * Creates a GUI
	 * @memberof HarvestListModule
	 * @class DecisionManager
	 * @classdesc It decides what to do with the messages from the server
	 */
	var DecisionManager = function DecisionManager() {
		this.GUI = new GUI();
	};
	DecisionManager.prototype =
		/** @lends HarvestListModule.DecisionManager */
		{
			/**
			 * It is called when a message has arrived from the server. It decides what to call
			 * @param  {object} message The message from the server
			 */
			decide: function(message) {
				switch (message.Name) {
					case "repositories-status-list":
						this.changeAllStatus(message);
						break;
					case "repository-status-changed":
						this.changeSingleStatus(message);
						break;
				}
			},
			/**
			 * It changes the status of all the repositories
			 * @param  {object} message The message from the server
			 */
			changeAllStatus: function(message) {
				this.GUI.changeAllStatus(message.Value);
			},
			/**
			 * It changes the status for a single repository
			 * @param  {object} message The message from the server
			 */
			changeSingleStatus: function(message) {
				this.GUI.changeGUI({
					ID: message.Repository,
					Status: message.Value
				});
				this.GUI.activateActionListeners();
			}
		};
	/**
	 * It gets the JQUERY element
	 * @memberof HarvestListModule
	 * @class GUI
	 * @classdesc It encapsulates the UX functionality
	 */
	var GUI = function GUI() {
		this.list = $("#repositories-list");
	};
	GUI.prototype =
		/** @lends HarvestListModule.GUI */
		{
			/**
			 * It changes the status of all the repositories
			 * @param  {array} repositories The array with the repositories
			 */
			changeAllStatus: function(repositories) {
				var repository, index;
				for (index = 0; index < repositories.length; index++) {
					repository = repositories[index];
					this.changeGUI(repository);
				}
				this.activateActionListeners();
			},
			/**
			 * It changes the status and action for a single repository and activates the listeners for actions
			 * @param  {object} repository A reference to the object which contains the ID and the Status of the repository
			 */
			changeGUI: function(repository) {
				this.changeStatus(repository);
				this.changeAction(repository);
			},
			/**
			 * It changes the HTML status according to the repository status
			 * @param  {object} repository The repository object
			 */
			changeStatus: function (repository) {
				var htmlID = this.getHTMLID(repository.ID),
					htmlSpan = this.getStatusColor(repository.Status);
				this.list.find(htmlID).html(htmlSpan);
			},
			/**
			 * It changes the action according to the status
			 * @param  {object} repository The repository object
			 */
			changeAction: function (repository) {
				var htmlID = this.getHTMLAction_ID(repository.ID),
					htmlSpan = this.getAction(repository);
				this.list.find(htmlID).html(htmlSpan);
			},
			/**
			 * It activates the listeners for the status actions
			 */
			activateActionListeners: function() {
				wisply.repositoriesModule.GUI.activateActionListeners();
			},
			/**
			 * It gets the HTML span for a status
			 * @param  {string} status The status of a repository
			 * @return {string} The HTML span for a status
			 */
			getStatusColor: function(status) {
				return wisply.repositoriesModule.GUI.getStatusColor(status);
			},
			/**
			 * It returns the JQUERY ID of a repository
			 * @param  {number} id The id of the repository
			 * @return {string} The JQUERY ID of a repository
			 */
			getHTMLID: function(id) {
				return this.getID("status", id);
			},
			/**
			 * It returns the ID of the action HTML element
			 * @param  {number} id The ID of the repository
			 */
			getHTMLAction_ID: function(id) {
				return this.getID("action", id);
			},
			/**
			 * Returns the HTML id for an action
			 * @param  {string} name The name of action
			 * @param  {number} id   The ID of the repository
			 */
			getID: function (name, id) {
				return "#rep-" + name + "-" + id;
			},
			/** It returns the HTML code for the actions, based on the status
			 * @param  {object} repository The repository object
			 * @return {string} The HTML code for actions
			 */
			getAction: function(repository) {
				var action = "<span class='text-muted'>Working</span>";
				switch (repository.Status) {
					case "ok":
					case "unverified":
						action = "<a href=''> <span data-toggle='tooltip' data-ID=" + repository.ID + " data-placement='top' title='' data-original-title='Update' class='repositories-init-harvest glyphicon glyphicon-retweet hover' ></span></a>";
						break;
					case "problems":
					case "verification-failed":
						action = "<a href='/admin/log/'>See log</a>";
					break;
				}
				return action;
			},
		};
	return {
		DecisionManager: DecisionManager,
		Stages: Stages
	};
};
$(document).ready(function() {
	"use strict";
	var module = new HarvestListModule();
	wisply.loadModule("harvest-list", module);
});
