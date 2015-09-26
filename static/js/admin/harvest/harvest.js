/* globals $, Websockets, server*/
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Harvest
 */
var harvestHistory = {};
var Harvest = function () {
	'use strict';
	/**
	 * Creates an empty history
	 * @memberof Harvest
	 * @class History
	 * @classdesc It holds a history of events
	 */
	var History = function History(callbackUpdate) {
		this.data = [];
		this.callbackUpdate = callbackUpdate;
	};
	History.prototype =
		/** @lends Harvest.History */
		{
			/**
			 * It adds a log message
			 * @param [string] content The content of the event
			 */
			log: function (content) {
				this.add(content, "LOG");
			},
			/**
			 * It adds an error message
			 * @param [string] content The content of the message
			 */
			logError: function (content) {
				this.add(content, "ERROR");
			},
			/**
			 * It adds a warning message
			 * @param [string] content The content of the message
			 */
			logWarning: function (content) {
				this.add(content, "WARN");
			},
			/**
			 * It is called by log, logError and logWarning. It adds an event to history. It assigns the datestamp. It calls the updater
			 * @private
			 * @param {string} content The content of the message
			 * @param {string} type    The type of the message. It can be "LOG", "ERROR" or "WARN"
			 */
			add: function (content, type) {
				var datetime;
				/**
				 * It returns the date in a human readable form
				 * @return {string} The date in a human readable form
				 */
				function getHumanDate() {
					var currentdate = new Date(),
						datetime = currentdate.getDate() + "." + (currentdate.getMonth() + 1) + "." + currentdate.getFullYear() + " at " + currentdate.getHours() + ":" + currentdate.getMinutes() + ":" + currentdate.getSeconds();
					return datetime;
				}
				datetime = getHumanDate();
				this.data.unshift({
					date: datetime,
					content: content,
					type: type
				});
				if (this.callbackUpdate) {
					this.callbackUpdate();
				}
			},
			/**
			 * It returns the history in HTML format
			 * @return [string] The history in HTML format
			 */
			getHTML: function () {
				/**
				 * It creates the HTML header of the table
				 * @return {string} The HTMl header of the table
				 */
				function getHeader() {
					var header = "";
					header = "<thead><tr><th class='text-center'>Date</th><th class='text-center'>Category</th><th class='text-center'>Content</th></tr></thead>";
					return header;
				}
				/**
				 * It creates the body of the table
				 * @param  {string} arrray The events
				 * @return {string}        The HTML body of the table
				 */
				function getBody(arrray) {
					var result = "<tbody>",
						i, currentEvent;
					/**
					 * It returns the type of HTML code
					 * @param  {string} type It can be "LOG", "ERROR" or "WARN"
					 * @return {string}      The HTML code for the type of the event
					 */
					function getType(type) {
						var textClass = "",
							content = "";
						switch (type) {
						case "LOG":
							textClass = "";
							content = "Event";
							break;
						case "ERROR":
							textClass = "text-danger";
							content = "Error";
							break;
						case "WARN":
							textClass = "text-warning";
							content = "Warning";
							break;
						}
						return "<span class='" + textClass + "'>" + content + "</span>";
					}
					for (i = 0; i < arrray.length; i++) {
						currentEvent = arrray[i];
						result += "<tr>";
						result += "<td>" + currentEvent.date + "</td>";
						result += "<td>" + getType(currentEvent.type) + "</td>";
						result += "<td>" + currentEvent.content + "</td>";
						result += "</tr>";
					}
					result += "</tbody>";
					return result;
				}
				var html = "<table class='table table=condensed table-hover ''>",
					events = "";
				html += getHeader();
				events = getBody(this.data);
				html += events;
				html += "</table>";
				return html;
			}
		};
	/**
	 * The constructor creates the history, page and stage manager
	 * @memberof Harvest
	 * @class Manager
	 * @classdesc It contains references to the Page object, StageManager and History
	 * @param [Repository] repository A reference to the repository
	 */
	var HarvestConnection = function HarvestConnection(manager) {
		var instance = this;
		this.manager = manager;
		var websockets = new Websockets(),
			host = server.host + "/admin/harvest/init/ws";
		this.connection = new websockets.Connection(host, this);
		var open = (function () {
			var conn = instance;
			return function () {
				conn.manager.onOpen();
			};
		})();
		var error = (function () {
			var conn = instance;
			return function () {
				conn.manager.onError();
			};
		})();
		this.onOpen = open;
		this.onClose = error;
		this.OnError = error;
	};
	HarvestConnection.prototype =
		/** @lends Harvest.Manager */
		{
			onMessage: function (evt) {
				var message = JSON.parse(evt.data),
					description = "";
				/**
				 * It processes the messages received from the server
				 * @param [event] evt The event which has been generated
				 */
				function createHumanDescription(msg) {
					/**
					 * It returns a description of the content of the message
					 * @param  {object} content The content of the message
					 * @return {string} The description of the content
					 */
					function getContentMessage(content) {
						if (content) {
							return " with content [" + content + "]";
						}
						return ", which does not has content";
					}
					/**
					 * It returns a human-readable message for the id of the repository
					 * @param  {number} id The id of the repository
					 * @return {string} Human readable message
					 */
					function getRepo(current, id) {
						if (current === id) {
							return "this repository";
						}
						return "the repository number " + id;
					}
					return ("I received from server the socket <b>" + msg.Name + "</b>" + getContentMessage(msg.Content) + " for " + getRepo(msg.Repository.id, msg.Repository) + ".");
				}
				description = createHumanDescription(message);
				harvestHistory.log(description);
				this.manager.onMessage(message);
			},
			send: function (msg) {
				this.connection.send(msg);
			},
		};
	/**
	 * The constructor creates the history, page and stage manager
	 * @memberof Harvest
	 * @class Manager
	 * @classdesc It contains references to the Page object, StageManager and History
	 * @param [Repository] repository A reference to the repository
	 */
	var Manager = function Manager(stageManager, decisionManager) {
		stageManager.manager = this;
		decisionManager.manager = this;
		// cross reference
		stageManager.decision = decisionManager;
		decisionManager.stage = stageManager;
		this.stageManager = stageManager;
		this.decisionManager = decisionManager;
	};
	Manager.prototype =
		/** @lends Harvest.Manager */
		{
      start: function () {
        this.connection = new HarvestConnection(this);
      },
      onOpen: function () {
        console.log("Start live :) ");
        this.stageManager.start();
      },
			onClose: function () {
				this.stageManager.stop();
			},
			onMessage: function (message) {
				this.decisionManager.decide(message);
			},
			send: function (msg) {
				this.connection.send(msg);
			},
		};
	/**
	 * The constructor creates the history, page and stage manager
	 * @memberof Harvest
	 * @class Manager
	 * @classdesc It contains references to the Page object, StageManager and History
	 * @param [Repository] repository A reference to the repository
	 */
	var StageManager = function StageManager(stages) {
		this.status = "stopped";
		this.currentStage = undefined;
		this.stages = stages;
		this.startingStageId = 0;
		// stages
	};
	StageManager.prototype =
		/** @lends Harvest.Manager */
		{
			/**
			 * It starts the manager. It calls the first stage
			 */
			setGUI: function (GUI) {
				this.GUI = GUI;
			},
			start: function () {
				this.performStage(this.startingStageId);
			},
      /**
       * It sends a message to the server. It can be modified to send a custom message
       * @param  {string} name  The name of the message
       * @param  {object} value The value of the message
       */
      sendMessage: function(name, value) {
        var msg = {
          Name: name,
          Value: value
        };
        this._send(msg);
      },
      /**
       * It sends a message to the server. YOU ARE NOT ALLOWED TO MODIFY
       * @private
       * @param  {object} value The object of the message
       */
      _send: function(msg) {
        this.manager.send(msg);
      },
			/**
			 * It performs a stage
			 * @param  {number} id The id of the stage
			 */
			performStage: function (ID) {
				var stage = this.stages[ID];
				harvestHistory.log("Starting stage <b>" + stage.name + "</b>...");
				this.status = "running";
				this.currentStage = stage;
				this.updateGUI();
				stage.perform(this);
			},
			/**
			 * It calls the next stage. If there are no stages, it calls firedEnd
			 */
			next: function () {
				var nextStageID = this.currentStage.id + 1;
				if (nextStageID >= this.stages.length) {
					this.firedEnd();
				} else {
					this.performStage(nextStageID);
				}
			},
			/**
			 * It forces the manager to stop. It forces the current stage to stop
			 */
			forceStop: function () {
				harvestHistory.log("The stage manager has been forced to stop.");
				this.current = "Stopped";
				this.state = "stopped";
				if (this.stage.stop) {
					this.stage.stop();
				}
			},
			/**
			 * Called when a stage has finished. It updates the page and calls the next stage
			 */
			firedStageFinished: function () {
				if (this.state === "stopped" || this.state === "paused") {
					harvestHistory.log("Imposible to continue!");
				} else {
					harvestHistory.log("Stage " + (this.current + 1) + " finished!");
					this.next();
				}
			},
			/**
			 * It is called when all the stages has been called. It updates the page
			 */
			firedEnd: function () {
				this.status = "finish";
				this.updateGUI();
				harvestHistory.log("The process has been finished!");
			},
			/**
			 * It performs again a stage
			 * @param  {number} number The id of the stage
			 */
			restart: function (number) {
				harvestHistory.log("Restarting from stage " + (number + 1) + "...");

				if(this.GUI) {
					this.GUI.restart();
				}
				this.performStage(number);
			},
			/**
			 * It pauses the manager
			 */
			pause: function () {
				this.status = "paused";
				this.updateGUI();
			},
			updateGUI: function () {
				if (this.GUI) {
					this.GUI.update();
				}
			},
      getCurrentStageID: function() {
          if(this.status ==="finished") {
            return this.stages.length;
          }
          if(!this.currentStage) {
            return "NOT STARTED";
          }
          return this.currentStage.id;
      },
			getCurrentProcent: function () {
				var current = this.getCurrentStageID(),
        percent, total;

        if(current === "NOT STARTED") {
          percent = 0;
        } else {
					total = this.stages.length;
					percent = 0;
				percent = (current) / total * 100;
      }
				return percent;
			}
		};
	return {
		Manager: Manager,
		History: History,
		StageManager: StageManager,
	};
};
$(document).ready(function () {
	"use strict";
	var module = new Harvest();
	harvestHistory = new module.History();
});
