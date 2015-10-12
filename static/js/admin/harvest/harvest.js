/* globals $, Websockets, server, wisply */
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
	 */
	var History = function History() {
		this.data = [];
		this.start = 0;
		this.eventsOnPage = 10;
		this.see = this.eventsOnPage;
	};
	History.prototype =
		/** @lends Harvest.History */
		{
			/**
			 * It logs an event
			 * @param [Harvest.HistoryGUI] gui The reference to the gui
			 */
			setGUI: function(id) {
				this.gui = new HistoryGUI(id, this);
			},
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
				if (this.gui) {
					this.gui.refresh();
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
					header = "<thead><tr><th class='text-center'>Number</th><th class='text-center'>Date</th><th class='text-center'>Category</th><th class='text-center'>Content</th></tr></thead>";
					return header;
				}
				var events = "<tbody>",
					moreButton = "",
					index, currentEvent;
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
				var dif = this.data.length - this.see,
					see = 0;
				if (dif <= 0) {
					see = this.data.length - 1;
				} else {
					see = this.see;
					moreButton = "You have seen " + this.see + " events. There are " + dif + " events more <br /><div class='text-center'><button id='harvest-history-see-more' class='btn btn-info'>See more</button></div>";
				}
				for (index = this.start; index <= see; index++) {
					currentEvent = this.data[index];
					events += "<tr>";
					events += "<td>" + (this.data.length - index) + "</td>";
					events += "<td>" + currentEvent.date + "</td>";
					events += "<td>" + getType(currentEvent.type) + "</td>";
					events += "<td>" + currentEvent.content + "</td>";
					events += "</tr>";
				}
				events += "</tbody>";
				var html = "<table class='table table=condensed table-hover ''>";
				html += getHeader();
				html += events;
				html += "</table>";
				html += moreButton;
				return html;
			},
			/**
			 * It allows the user to see more events
			 */
			seeMore: function() {
				this.see = this.see + this.eventsOnPage;
				this.gui.refresh();
			},
			/**
			 * It sets the start and end to the default values
			 */
			reset: function() {
				this.see = this.eventsOnPage;
				this.start = 0;
			}
		};
	/**
	 * The constructor creates the connection
	 * @memberof Harvest
	 * @class HarvestConnection
	 * @classdesc It is a middleware between the harvest manager and the web sockets connection
	 * @param [string] id The id of the element where the history will be displayed
	 * @param [Harvest.History] history The reference to the gui
	 */
	var HistoryGUI = function HistoryGUI(id, history) {
		this.activ = false;
		this.element = $(id);
		this.history = history;
	};
	HistoryGUI.prototype =
		/** @lends Harvest.HistoryGUI */
		{
			/**
			 * It updates the gui
			 */
			refresh: function() {
				if (this.activ) {
					this.element.html(this.history.getHTML());
					$("#harvest-history-see-more").click(function() {
						harvestHistory.seeMore();
					});
				}
			},
			/**
			 * It activates the update of gui
			 */
			activate: function() {
				this.history.reset();
				this.activ = true;
				this.refresh();
			},
			/**
			 * It disables the gui
			 */
			disable: function() {
				this.activ = false;
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
						} else if (id === 0) {
							return "all repositories";
						}
						return "the repository number " + id;
					}
					/**
					 * It describes some actions
					 * @param  {string} name The name of the message
					 */
					function getName(message) {
						var toReturn = "";
						toReturn += "<strong>" + message.Name + "</strong>";
						if (message.Name === "repository-status-changed") {
							toReturn += " " + wisply.repositoriesModule.GUI.getStatusColor(message.Value.trim());
						}
						return toReturn;
					}
					return ("I received from server the socket " + getName(msg) + " " + getContentMessage(msg.Value) + " for " + getRepo(msg.Repository.id, msg.Repository) + ".");
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
					harvestHistory.log("Stage <b>" + this.currentStage.name + "</b> finished!");
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
