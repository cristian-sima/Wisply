/* global $, Harvest, wisply, server, harvestHistory, CountUp */
/**
 * @file Encapsulates the functionality for the harvest process
 * @author Cristian Sima
 */
/**
 * @namespace HarvestProcess
 */
var HarvestProcess = function() {
	'use strict';
	var Stages = [{
		id: 0,
		name: "Prepare...",
		/**
		 * It modifies the sendMessage method of stage manager and load the current repository
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			var repository = {};
			// need to store once
			this.manager = manager;
			this.manager.GUI.start();
			// modify stage manager send function
			// this is to include everytime the id of the current repository
			this.manager.sendMessage = function(name, value) {
				var msg = {
					Name: name,
					Value: value,
					Repository: this.repository.id
				};
				this._send(msg);
			};
			// load Repository
			repository = new wisply.repositoriesModule.Repository(server.repository);
			// load Repository
			this.manager.repository = repository;
			this.manager.GUI.updateRepositoryStatus();
			setTimeout(function() {
				manager.firedStageFinished();
			}, 1000);
		}
	}, {
		id: 1,
		name: "Getting information...",
		/**
		 * It tells the server to send the process' details regarding this repository
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			this.manager = manager;
			manager.sendMessage("get-current-progress", "");
		},
		/**
		 * It analyses the message from the server
		 * @param  {object} message The message from the server
		 */
		analyse: function(message) {
			if (message.Value !== null) {
				window.location="/admin/repositories";
			} else {
				harvestHistory.log("We start a new process");
				this.initNewProcess();
			}
		},
		/**
		 * It is called when the process exists on the server. It sends on the current stage
		 * @param  {object} contents The value of the message
		 */
		initExistingProcess: function(serverManager) {
			harvestHistory.log("Load process from stage <b>" + this.manager.stages[serverManager.CurrentAction].name + "</b>.");
			this.manager.repository.Identification = serverManager.Identification;
			this.manager.existingProcessActions = serverManager.Actions;
			this.manager.performStage(serverManager.CurrentAction);
		},
		/**
		 * It is called when there is no process on the server. It calls the next stage
		 */
		initNewProcess: function() {
			this.manager.firedStageFinished();
		}
	}, {
		id: 2,
		name: "Start process...",
		/**
		 * It tells the server to start the init process
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			this.manager = manager;
			this.paint();
			this.manager.sendMessage("start-progress", "");
		},
		/**
		 * It shows the loading image
		 */
		paint: function() {
			this.manager.GUI.showCurrent(wisply.getLoadingImage("big"));
		}
	}, {
		id: 3,
		name: "Validation...",
		/**
		 * It saves the reference of manager
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			this.manager = manager;
		},
		/**
		 * It disables the possibility to modify the URL
		 */
		disableModifyURL: function() {
			$('#modifyButton').prop('disabled', true);
			$('#Source-URL').prop('disabled', true);
		},
		/**
		 * It enables the possibility to modify the URL
		 */
		enableModifyURL: function() {
			var field = $('#Source-URL');
			$('#modifyButton').prop('disabled', false);
			field.prop('disabled', false);
			field.focus();
		},
		/**
		 * It tells the server to change the base URL
		 * @param  {string} newURL The new base URL for repository
		 */
		changeURL: function(newURL) {
			this.manager.sendMessage("change-url", newURL);
		}
	}, {
		id: 4,
		name: "Identifying...",
		/**
		 * It saves the reference of manager
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			this.manager = manager;
			this.init();
		},
		init: function() {
			$("#URL-input").toggle();
			$("#Name-Repository").toggle();
			$("#modifyButton").hide();
		},
		/**
		 * It remembers the identification details, paints them and go to the next stage
		 * @param  {object} indentification The object which holds all the identification details
		 */
		analyse: function(identification) {
			this.manager.repository.Identification = identification;
			this.manager.firedStageFinished();
		}
	}, {
		id: 5,
		name: "Harvesting...",
		lastTime: 0,
		/**
		 * It saves the reference of manager
		 * @param  {Harvest.StageManager} manager The reference to the stage manager
		 */
		perform: function(manager) {
			this.manager = manager;
			this.counters = {};
			this.init();
		},
		/**
		 * It adds the counters
		 */
		init: function() {
			if (this.manager.existingProcessActions) {
				this.manager.stages[4].perform(this.manager);
			}
			this.showCounters();
		},
		showCounters: function() {
			this.counters = [
				new WisplyCounter({
					name: "Formats",
					type: "formats",
				}),
				new WisplyCounter({
					name: "Collections",
					type: "collections",
				}),
				new WisplyCounter({
					name: "Records",
					type: "records",
				}),
				new WisplyCounter({
					name: "Identifiers",
					type: "identifiers",
				}),
			];
			this.manager.GUI.showCurrent('<div id="harvesting-counters"></div>');
			this.manager.GUI.drawCounters(this.counters, "harvesting-counters");
		},
		/**
		 * It sets the current counter
		 * @param  {string} type The name of the counter
		 */
		setCurrentCounter: function(type) {
			this.currentCounter = this.getCounter(type);
		},
		/**
		 * It analyse the result from the server and decides which counter to update
		 * @param  {object} result The message from the server
		 */
		analyse: function(result) {
				var now;

				harvestHistory.log("Server told me that Wisply harvested " + result.Number + " " + result.Operation);
				if(!this.currentCounter || this.currentCounter.type != result.Operation) {
					console.log("Start");
					this.lastTime = result.Number;
					this.setCurrentCounter(result.Operation);
					this.currentCounter.start(result.Number);
				}
				else {
					now = this.lastTime + parseInt(result.Number, 10);
				  this.getCounter(result.Operation).update(now);
					this.lastTime = now;
				}
		},
		/**
		 * It updates the counter
		 * @param  {string} requested  The type of counter
		 */
		getCounter: function(requested) {
			var i, counter;
			for (i = 0; i < this.counters.length; i++) {
				counter = this.counters[i];
				if (counter.type === requested.toLowerCase()) {
					return counter;
				}
			}
			throw "There is no counter with the name " + requested;
		},
		/**
		 * It is called by stage manager. In case there is a current counter, it tells it to show the error
		 */
		stop: function() {
			if (this.currentCounter) {
				this.currentCounter.showError();
			}
		},
		/**
		 * It is called when the stage has finished. It deletes the counters, the current counter and goes to the next stage
		 */
		end: function() {
			delete this.currentCounter;
			delete this.counters;
			this.manager.firedStageFinished();
		}
	}];
	/**
	 * The constructor does nothing
	 * @memberof HarvestProcess
	 * @class WisplyCounter
	 * @classdesc It decides what to do with the messages from the server
	 */
	var WisplyCounter = function WisplyCounter(info) {
		this.type = info.type;
		this.name = info.name;
		this.object = undefined;
		this.stopped = false;
	};
	WisplyCounter.prototype =
		/** @lends HarvestProcess.WisplyCounter */
		{
			/**
			 * It updates the counter
			 */
			update: function(value) {
				this.object.update(value);
			},
			/**
			 * It sets the counter yellow and starts it
			 */
			start: function(value) {
				this._getElements();
				this.container.removeClass("text-muted");
				this.element.addClass("label-warning");
				this.element.css({
					"color": "#ffffff"
				});
				this._start(value);
			},
			/**
			 * It gets the JQuery elements. It is private
			 * @ignore
			 */
			_getElements: function() {
				this.container = $("#repository-counter-container-" + this.type);
				this.element = $("#repository-counter-" + this.type);
			},
			/**
			 * It creates the counter and starts the counter. It is private
			 * @param [number] value The value to go to
			 * @ignore
			 */
			_start: function(value) {
				var instance = this,
					options = {  
						useEasing: false,
						  useGrouping: true,
						  separator: ',',
						  decimal: '.',
						  prefix: '',
						  suffix: ''
					},
					time = (value / 80);
				if (time > 10) {
					time = 10;
				}
				this.object = new CountUp("repository-counter-" + this.type, 0, value, 0, time, options);
				this.object.start(function() {
					instance._finish();
				});
			},
			/**
			 * It is called when the counter has finished
			 * @ignore
			 */
			_finish: function() {
				if (this.stopped) {
					this.element.removeClass("label-warning");
					this.element.addClass("label-success");
				}
			},
			/**
			 * It sets the stopped value to true. In case the counter has finished, it calls _finish
			 */
			finish: function() {
				this.stopped = true;
				if (this.object.remaining <= 0) {
					this._finish();
				}
			},
			/**
			 * It makes the counter red
			 */
			showError: function() {
				this.element.removeClass("text-muted label-warning label-success");
				this.element.addClass("label-danger");
			},
			getHTML: function() {
				return '<div id="repository-counter-container-' + this.type + '" class="text-center text-muted col-md-3">' + "<div class='counter-name'>" + this.name + '</div>' + '<span class="label label-as-badge big-number text-muted" id="repository-counter-' + this.type + '">0</span>' + '</div>';
			}
		};
	/**
	 * The constructor does nothing
	 * @memberof HarvestProcess
	 * @class DecisionManager
	 * @classdesc It decides what to do with the messages from the server
	 */
	var DecisionManager = function DecisionManager() {};
	DecisionManager.prototype =
		/** @lends HarvestProcess.DecisionManager */
		{
			/**
			 * It is called when a message has arrived from the server. It decides what to call
			 * @param  {object} message The message from the server
			 */
			decide: function(message) {
				if (this.isGoogMessage(message)) {
					switch (message.Name) {
						case "repository-status-changed":
							this.stage.repository.status = message.Value;
							this.stage.GUI.updateRepositoryStatus();
							switch(message.Value) {
								case "verifying":
									this.stage.performStage(3);
								break;
									case "initializing":
									this.stage.firedStageFinished();
									this.stage.firedStageFinished();
								break;
								case "verification-failed":
									this.stage.GUI.showCurrent("The verification failed");
									this.stage.pause();
									this.stage.stages[3].enableModifyURL();
								break;
								case "verified":
									this.stage.firedStageFinished();
								break;
							}
							break;
						case "harvest-update":
							this.stage.stages[5].analyse(message.Value);
						break;
						case "existing-process-on-server":
							this.stage.currentStage.analyse(message);
							break;
						case "process-finished":
							this.stage.firedStageFinished();
						break;
					}
				}
			},
			/**
			 * It checks that a message is for this process. It is if it contains a "Repository" field and the "Repository" field is the same as the current repository ID
			 * @param  {object} message The message from server
			 */
			isGoogMessage: function(message) {
				return (message.Repository) && (message.Repository === this.stage.repository.id);
			}
		};
	/**
	 * Gets the JQuery elements and creates the indicator
	 * @memberof HarvestProcess
	 * @class StageGUI
	 * @classdesc It contains the functionality for GUI
	 */
	var StageGUI = function StageGUI(manager) {
		this.manager = manager;
		this.element = $("#stages");
		this.current = $("#current");
		this.indicator = new Indicator();
	};
	StageGUI.prototype =
		/** @lends HarvestProcess.StageGUI */
		{
			/**
			 * It shows the stage, calls update and load listeners
			 */
			start: function() {
				this.element.slideDown();
				this.update();
				this.loadListeners();
			},
			/**
			 * It activates the listeners
			 */
			loadListeners: function() {
				var instance = this;
				$("#modifyButton").click(function() {
					instance.manager.stages[3].changeURL($("#Source-URL").val());
					instance.manager.stages[3].disableModifyURL();
					instance.manager.restart(2);
				});
				harvestHistory.setGUI("#history");
				$("#historyButton").click(function() {
					harvestHistory.gui.activate();
				});
				$("#currentButton").click(function() {
					harvestHistory.gui.disable();
				});
			},
			/**
			 * It updates the list, the indicator and the process status
			 */
			update: function() {
				this.updateList();
				this.updateIndicator();
				this.updateProcessStatus();
			},
			/**
			 * It modify the HTML of current panel
			 * @param [string] html The HTML code to be inserted
			 */
			showCurrent: function(html) {
				this.current.html(html);
			},
			/**
			 * It updates the list of current stages
			 */
			updateList: function() {
				var container = this.element.find("#stage-list"),
					manager = this.manager,
					current = manager.getCurrentStageID(),
					stages = manager.stages,
					stage, id, html = "",
					item = "";
				for (id = 0; id < stages.length; id++) {
					item = "";
					stage = stages[id];
					if (id === current) {
						item += '<li class="list-group-item active">';
						item += stage.name;
						item += "</li>";
					} else {
						if (id < current) {
							item = '<li class="list-group-item text-muted"><del>' + stage.name + '</del></li>';
						} else {
							item = '<li class="list-group-item">' + stage.name + "</li>";
						}
					}
					html += item;
				}
				if (current === stages.length) {
					html += '<div class="panel panel-success">  <div class="panel-heading">    <h3 class="panel-title">Done!</h3></div>  <div class="panel-body">    The process is over.  </div></div>';
				}
				container.html(html);
			},
			/**
			 * It updates the indicator
			 */
			updateIndicator: function() {
				var procent;
				if (this.manager.status === "finish") {
					this.indicator.finished();
				}
				procent = this.manager.getCurrentProcent();
				this.indicator.set(procent);
			},
			/**
			 * It updates the general status of the process
			 */
			updateRepositoryStatus: function() {
				var status = this.manager.repository.status,
					html = "Status: ",
					span = "";
				span = wisply.repositoriesModule.GUI.getStatusColor(status);
				html += span;
				$("#repository-status").html(html);
			},
			/**
			 * It tells the indicator to start
			 */
			restart: function() {
				this.indicator.start();
			},
			/**
			 * It tells the indicator to change to yellow and stop it
			 */
			pause: function() {
				this.indicator.warning();
				this.indicator.stop();
			},
			/**
			 * It tells the indicator to change to red and stop it
			 */
			stop: function() {
				this.indicator.error();
				this.indicator.stop();
			},
			/**
			 * It draws the Wisply counters in the element
			 * @param  {array} counters An array with WisplyCounters
			 * @param  {string} id       The id of the html element
			 */
			drawCounters: function(counters, id) {
				var text = "",
					html = "",
					counter = {},
					i;
				for (i = 0; i < counters.length; i++) {
					counter = counters[i];
					html += counter.getHTML();
				}
				text += '<div class="row text-center" id="repository-elements" style="display:block">';
				text += html;
				text += "</div>";
				$("#" + id).html(text);
			},
			/**
			 * It updates the general status of the process
			 */
			updateProcessStatus: function() {
				var status = this.manager.status,
					html = "Progress: ";
				switch (status) {
					case "stopped":
						html += '<span class="label label-danger">Stopped</span>';
						this.stop();
						break;
					case "paused":
						html += '<span class="label label-warning">Paused</span>';
						this.pause();
						break;
					case "finish":
						html += '<span class="label label-success">Finish</span>';
						break;
					case "running":
						html += wisply.getLoadingImage("small");
						break;
					default:
						html += "Problem";
						break;
				}
				$("#process-status").html(html);
			}
		};
	/**
	 * Gets the JQuery element
	 * @memberof HarvestProcess
	 * @class Indicator
	 * @classdesc It represents a visual indicator for the current progress of the process
	 */
	var Indicator = function Indicator(gui) {
		this.gui = gui;
		this.element = $("#general-indicator");
		this.animation = undefined;
	};
	Indicator.prototype =
		/** @lends HarvestProcess.Indicator */
		{
			/**
			 * It sets the indicator to a percent
			 * @param  {number} percent The percent to set
			 */
			set: function(percent) {
				if (this.animation) {
					this.animation.finish();
				}
				this.animation = this.element.find(".progress-bar").animate({
					"width": percent + "%"
				}, 100);
			},
			/**
			 * It removes the stipes (it makes it static)
			 */
			stop: function() {
				this.element.removeClass("progress-striped");
			},
			/**
			 * It makes it default and adds the stipes
			 */
			start: function() {
				this.element.find(".progress-bar").removeClass("progress-bar-warning");
				this.element.find(".progress-bar").addClass("progress-bar-default");
				this.element.addClass("progress-striped");
			},
			/**
			 * It change the colour to green
			 */
			finished: function() {
				this.changeIndicator("success");
			},
			/**
			 * It changes the design of indicator for a warning situation
			 */
			warning: function() {
				this.changeIndicator("warning");
			},
			/**
			 * It changes the design of indicator for an error situation
			 */
			error: function() {
				this.changeIndicator("danger");
			},
			/**
			 * It changes the design of the indicator for a certain situation
			 * @param  {string} type The situation. It can be "danger" or "warning" or "success"
			 */
			changeIndicator: function(type) {
				var element = this.element.find(".progress-bar");
				element.removeClass();
				element.addClass("progress-bar-" + type);
				element.addClass("progress-bar");
				this.stop();
			},
		};
	return {
		DecisionManager: DecisionManager,
		Stages: Stages,
		StageGUI: StageGUI
	};
};
$(document).ready(function() {
	"use strict";
	var harvest,
		process,
		repository,
		decision,
		stage,
		manager,
		stages;
	harvest = new Harvest();
	process = new HarvestProcess();
	repository = wisply.repositriesModule;
	decision = new process.DecisionManager();
	stages = process.Stages;
	stage = new harvest.StageManager(stages);
	stage.setGUI(new process.StageGUI(stage));
	manager = new harvest.Manager(stage, decision);
	wisply.manager = manager;
	manager.start();
});
