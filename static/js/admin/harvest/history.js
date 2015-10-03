/* global $, Harvest, wisply*/
/**
 * @file Encapsulates the functionality for list harvesting
 * @author Cristian Sima
 */
/**
 * @namespace HarvestList
 * It holds the functionality to see a live list of repositories
 */
var HistoryEvent = function() {
	'use strict';

	var TimeDifference = {
		seconds: 0,
		minutes: 0,
		hours: 0,
		days: 0,
	};

	var EventHistory = function EventHistory() {

	};

	/**
	 * It creates a new date from a string
	 * @param {string} The date in a string
	 * @return {Date} The date in a Date object
	 * @see http://stackoverflow.com/questions/3075577/convert-mysql-datetime-stamp-into-javascripts-date-format
	 */
	function newDateFromString() {
		var t = "2010-06-09 13:12:01".split(/[- :]/);

		return new Date(t[0], t[1]-1, t[2], t[3], t[4], t[5]);
	}

	/**
	 * Calculates the time difference between 2 Javascript date objects
	 * @param  {Date} start The start date
	 * @param  {Date} end  The end date
	 * @return {HistoryEvent.TimeDifference} An object which contains the time difference
	 */
	function getDifference(start, end)
	{
			var total, difference,
				MILISECOND = 1,
				SECOND = 1000 * MILISECOND,
				MINUTE = 60 * SECOND,
				HOUR = MINUTE * 60,
				DAY = HOUR * 24;
			/**
			 * It returns the number of days given the number of milliseconds
			 * @param  {number} seconds The number of milliseconds
			 * @return {number} The number of days
			 */
			function getDays(seconds) {
				return Math.floor(total / DAY);
			}

			/**
			 * It returns the number of hours given the number of milliseconds
			 * @param  {number} seconds The number of milliseconds
			 * @return {number} The number of hours
			 */
			function getHours(seconds) {
				return Math.floor(total / HOUR);
			}

			/**
			 * It returns the number of minutes given the number of milliseconds
			 * @param  {number} seconds The number of milliseconds
			 * @return {number} The number of minutes
			 */
			function getMinutes(seconds) {
				return Math.floor(total / MINUTE);
			}

			/**
			 * It returns the number of minutes given the number of milliseconds
			 * @param  {number} seconds The number of milliseconds
			 * @return {number} The number of seconds
			 */
			function getSeconds(seconds) {
				return Math.floor(total / SECOND);
			}
      total = end.getTime() - start.getTime();
     	difference = new TimeDifference();
			difference.seconds = getSeconds(seconds);
			difference.minutes = getSeconds(minutes);
			difference.hours = getSeconds(hours);
			difference.days = getSeconds(days);
	}

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
			$("#harvest-history-container").on('hidden.bs.modal', function () {
				console.log("inchis");
					harvestHistory.gui.disable();
			});
		}
	}];
	/**
	 * Creates a GUI
	 * @memberof HarvestList
	 * @class DecisionManager
	 * @classdesc It decides what to do with the messages from the server
	 */
	var DecisionManager = function DecisionManager() {
		this.GUI = new GUI();
	};
	DecisionManager.prototype =
		/** @lends HarvestList.DecisionManager */
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
					case "status-changed":
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
				this.GUI.changeStatus({
					ID: message.Repository,
					Status: message.Value
				});
				this.GUI.activateActionListeners();
			}
		};
	/**
	 * It gets the JQUERY element
	 * @memberof HarvestList
	 * @class GUI
	 * @classdesc It encapsulates the UX functionality
	 */
	var GUI = function GUI() {
		this.list = $("#repositories-list");
	};
	GUI.prototype =
		/** @lends HarvestList.GUI */
		{
			/**
			 * It changes the status of all the repositories
			 * @param  {array} repositories The array with the repositories
			 */
			changeAllStatus: function(repositories) {
				var repository, index;
				for (index = 0; index < repositories.length; index++) {
					repository = repositories[index];
					this.changeStatus(repository);
				}
				this.activateActionListeners();
			},
			/**
			 * It changes the status for a single repository and activates the listeners for actions
			 * @param  {object} repository A reference to the object which contains the ID and the Status of the repository
			 */
			changeStatus: function(repository) {
				var htmlID = this.getHTMLID(repository.ID),
					htmlSpan = this.getStatusColor(repository.Status),
					action = this.getAction(repository);
				this.list.find(htmlID).html(htmlSpan + action);
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
			 * It returns the HTML code for the actions, based on the status
			 * @param  {object} repository The repository object
			 * @return {string} The HTML code for actions
			 */
			getAction: function(repository) {
				var action = "";
				switch (repository.Status) {
					case "unverified":
						action = "<a href=''> <span data-toggle='tooltip' data-ID='" + repository.ID + "' data-placement='top' title='' data-original-title='Start now!' class='repositories-init-harvest glyphicon glyphicon-sort-by-attributes hover' ></span></a>";
						break;
					case "verification-failed":
						action = "<a href=''> <span data-toggle='tooltip' data-ID='" + repository.ID + "' data-placement='top' title='' data-original-title='Try again' class='repositories-init-harvest glyphicon glyphicon glyphicon-refresh hover' ></span></a>";
				}
				return action;
			},
			/**
			 * It returns the JQUERY ID of a repository
			 * @param  {number} id The id of the repository
			 * @return {string} The JQUERY ID of a repository
			 */
			getHTMLID: function(id) {
				return "#rep-status-" + id;
			}
		};
	return {
		DecisionManager: DecisionManager,
		Stages: Stages
	};
};
$(document).ready(function() {
	"use strict";
	var harvest,
		list,
		repository,
		decision,
		stage,
		manager,
		stages;
	harvest = new Harvest();
	list = new HarvestList();
	repository = wisply.repositriesModule;
	decision = new list.DecisionManager();
	stages = list.Stages;
	stage = new harvest.StageManager(stages);
	manager = new harvest.Manager(stage, decision);
	wisply.manager = manager;
	manager.start();
});