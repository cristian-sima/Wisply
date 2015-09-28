/* globals $, Websockets, server*/
/**
 * @file Encapsulates the functionality for harvest process.
 * @author Cristian Sima
 */
// It has a history
var harvestHistory = {};
/**
 * @namespace Harvest
 */
var Harvest = function() {
	'use strict';
	/**
	 * Creates an empty history
	 * @memberof Harvest
	 * @class History
	 * @classdesc It holds a history of events
	 * @param {function} The function which is called when an event is added to history
	 */
	var History = function History(callbackUpdate) {
		this.data = [];
		this.callbackUpdate = callbackUpdate;
	};
	History.prototype =
		/** @lends Harvest.History */
		{
			/**
			 * It logs an event
			 * @param [string] content The content of the event
			 */
			log: function(content) {
				this.add(content, "LOG");
			},
			/**
			 * It adds an error message
			 * @param [string] content The content of the message
			 */
			logError: function(content) {
				this.add(content, "ERROR");
			},
			/**
			 * It adds a warning message
			 * @param [string] content The content of the message
			 */
			logWarning: function(content) {
				this.add(content, "WARN");
			},
			/**
			 * It is called by log, logError and logWarning. It adds an event to history. It assigns the datestamp. It calls the updater
			 * @private
			 * @param {string} content The content of the message
			 * @param {string} type    The type of the message. It can be "LOG", "ERROR" or "WARN"
			 */
			add: function(content, type) {
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
			getHTML: function() {
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
	 * The constructor creates the connection
	 * @memberof Harvest
	 * @class HarvestConnection
	 * @classdesc It is a middleware between the harvest manager and the web sockets connection
	 * @param [Repository] manager A reference to the harvest manager
	 */
	var HarvestConnection = function HarvestConnection(manager) {
		var instance = this;
		this.manager = manager;
		var websockets = new Websockets(),
			host = server.host + "/admin/harvest/init/ws";
		this.connection = new websockets.Connection(host, this);
		var open = (function() {
			var conn = instance;
			return function() {
				conn.manager.onOpen();
			};
		})();
		var error = (function() {
			var conn = instance;
			return function() {
				conn.manager.onClose();
			};
		})();
		this.onOpen = open;
		this.onClose = error;
		this.onError = error;
	};
	HarvestConnection.prototype =
		/** @lends Harvest.HarvestConnection */
		{
			/**
			 * It is called when a message arrived from server. It logs in the history and tells the manager
			 * @param  {object} evt The web socket event
			 */
			onMessage: function(evt) {
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
			/**
			 * It calls the websocket methdod to send the message
			 * @param  {object} message An object which encapsulates the message
			 */
			send: function(message) {
				this.connection.send(message);
			},
		};
	/**
	 * It saves the references
	 * @memberof Harvest
	 * @class HarvestConnection
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
			/**
			 * It creates the web socket connection
			 */
			start: function() {
				this.connection = new HarvestConnection(this);
			},
			/**
			 * It is called when the connection has been established. It starts the stage manager.
			 */
			onOpen: function() {
				console.log("Start live :) ");
				this.stageManager.start();
			},
			/**
			 * It is called when the ws connection is closed. It stops the stage manager
			 */
			onClose: function() {
				this.stageManager.stop();
			},
			/**
			 * It is called when a messsage has arrived from the server. It calls the decision method of decisionManager object
			 * @param  {object} message The message from the server
			 */
			onMessage: function(message) {
				this.decisionManager.decide(message);
			},
			/**
			 * It sends a message to the server using web sockets
			 * @param  {object} message The object which encapsulates the message
			 */
			send: function(msg) {
				this.connection.send(msg);
			},
		};
	/**
	 * Saves the stages
	 * @memberof Harvest
	 * @class StageManager
	 * @classdesc The stage manager encapsulates the functionality for managing the stages of the client
	 * @param [array] stages An array with the stages
	 */
	var StageManager = function StageManager(stages) {
		this.status = "stopped";
		this.currentStage = undefined;
		this.stages = stages;
		this.startingStageId = 0;
		// stages
	};
	StageManager.prototype =
		/** @lends Harvest.StageManager */
		{
			/**
			 * It sets a GUI for the manager. The GUI MUST implement theser methods: update, restart
			 * @param [object] GUI A reference to the GUI object
			 */
			setGUI: function(GUI) {
				this.GUI = GUI;
			},
			/**
			 * It starts the manager. It calls the first stage
			 */
			start: function() {
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
			 * It sends a message to the server. YOU ARE NOT ALLOWED TO MODIFY IT
			 * @private
			 * @param  {object} value The object of the message
			 */
			_send: function(msg) {
				this.manager.send(msg);
			},
			/**
			 * It performs a stage. Updates teh GUI and set the status "running"
			 * @param  {number} id The id of the stage
			 */
			performStage: function(ID) {
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
			next: function() {
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
			stop: function() {
				harvestHistory.log("The stage manager has been forced to stop.");
				this.current = "Stopped";
				this.status = "stopped";
				if (this.currentStage.stop) {
					this.currentStage.stop();
				}
				this.updateGUI();
			},
			/**
			 * Called when a stage has finished. It updates the page and calls the next stage
			 */
			firedStageFinished: function() {
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
			firedEnd: function() {
				this.status = "finish";
				this.updateGUI();
				harvestHistory.log("The process has been finished!");
			},
			/**
			 * It performs again a stage. If there is any GUI, it restarts it
			 * @param  {number} number The id of the stage
			 */
			restart: function(number) {
				harvestHistory.log("Restarting from stage " + (number + 1) + "...");
				if (this.GUI) {
					this.GUI.restart();
				}
				this.performStage(number);
			},
			/**
			 * It pauses the manager
			 */
			pause: function() {
				this.status = "paused";
				this.updateGUI();
			},
			/**
			 * It the manager has a GUI, it calls the method update of GUI
			 */
			updateGUI: function() {
				if (this.GUI) {
					this.GUI.update();
				}
			},
			/**
			 * It returns the id of current stage. If there is no current stage, it returns "NOT STARTED"
			 */
			getCurrentStageID: function() {
				if (this.status === "finish") {
					return this.stages.length;
				}
				if (!this.currentStage) {
					return "NOT STARTED";
				}
				return this.currentStage.id;
			},
			/**
			 * It returns the percent of actions done
			 * @return [number] The percent of actions done
			 */
			getCurrentProcent: function() {
				var current = this.getCurrentStageID(),
					percent, total;
				if (current === "NOT STARTED") {
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
$(document).ready(function() {
	"use strict";
	var module = new Harvest();
	harvestHistory = new module.History();
});
