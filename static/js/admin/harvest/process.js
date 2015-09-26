/* global $, Harvest, wisply*/
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Harvest
 */
var HarvestProcess = function() {
	'use strict';
	var Stages = [
		{
		id: 0,
		name: "Prepare...",
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
	},
	{
	id: 1,
	name: "Getting information...",
	perform: function(manager) {
		this.manager = manager;
		manager.sendMessage("getCurrentProcess", "");
	},
	analyse: function(message){
		if(message.Value !== null) {
				this.initExistingProcess(message.Value);
		} else {
			this.initNewProcess();
		}
	},
	initExistingProcess : function(contents) {
			this.manager.performStage(contents.CurrentAction);
	},
	initNewProcess: function () {
		this.manager.firedStageFinished();
	}
},
	{
		id: 2,
		name: "Start process...",
		perform: function(manager) {
			this.manager = manager;
			this.paint();
			this.manager.sendMessage("startInitializing", "");
		},
		paint: function() {
			this.manager.GUI.showCurrent(wisply.getLoadingImage("big"));
		}
	},
	{
		id: 3,
		name: "Validation...",
		perform: function(manager) {
			this.manager = manager;
		},
		/**
		 * It disables the possibility to modify the URL
		 */
		disableModifyURL: function () {
				$('#modifyButton').prop('disabled', true);
				$('#Source-URL').prop('disabled', true);
		},
		/**
		 * It enables the possibility to modify the URL
		 */
		enableModifyURL: function () {
				$('#modifyButton').prop('disabled', false);
				$('#Source-URL').prop('disabled', false);
		},
		changeURL: function(newURL) {
				this.manager.sendMessage("changeRepositoryURL", newURL);
		}
	},
	{
		id: 4,
		name: "Collecting records...",
		perform: function(manager) {
			this.manager = manager;
		}
	}
];
	var DecisionManager = function DecisionManager() {
		// this.GUI = new GUI();
	};
	DecisionManager.prototype =
		/** @lends Harvest.Manager */
		{
			decide: function(message) {
				console.log(message);
				if (this.isGoogMessage(message)) {
					switch (message.Name) {
						case "RepositoryChangedStatus":
							this.stage.repository.status = message.Value.NewStatus;
							this.stage.GUI.updateRepositoryStatus();
							if (message.Value.NewStatus === "verifying") {
								this.stage.performStage(3);
							}
							break;
						case "ProcessOnServer":
							this.stage.currentStage.analyse(message);
							break;
						case "VerificationFailed":
							this.stage.GUI.showCurrent(message.Value.Explication);
							this.stage.pause();
							this.stage.stages[3].enableModifyURL();
							break;
					}
				}
			},
			isGoogMessage: function(message) {
				return (message.Repository) && (message.Repository === this.stage.repository.id);
			}
		};
	var StageGUI = function StageGUI(manager) {
		this.manager = manager;
		this.element = $("#stages");
		this.current = $("#current");
		this.indicator = new Indicator();
	};
	StageGUI.prototype =
		/** @lends ListHarvest.GUI */
		{
			start: function () {
				this.element.slideDown();
				this.update();
				this.loadListeners();
			},
			loadListeners: function () {
				var instance = this;
				$("#modifyButton").click(function () {
						instance.manager.stages[3].changeURL($("#Source-URL").val());
						instance.manager.stages[3].disableModifyURL();
						instance.manager.restart(2);
				});
			},
			update: function() {
				this.updateList();
				this.updateIndicator();
				this.updateProcessStatus();
			},
			showCurrent: function(html) {
					this.current.html(html);
			},
			/**
			 * It updates the list of current stages
			 */
			updateList: function () {
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
			updateRepositoryStatus: function () {
					var status = this.manager.repository.status,
							html = "Status: ",
							span = "";

					span = wisply.repositoriesModule.GUI.getStatusColor(status);

					html += span;

					$("#repository-status").html(html);
			},
			restart : function () {
					this.indicator.start();
			},
			pause : function () {
					this.indicator.warning();
					this.indicator.stop();
			},
			/**
			 * It updates the general status of the process
			 */
			updateProcessStatus: function () {
				var status = this.manager.status,
						html = "Progress: ";
				switch (status) {
				case "stopped":
						html += '<span class="label label-danger">Stopped</span>';
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
	var Indicator = function Indicator(gui) {
		this.gui = gui;
		this.element = $("#general-indicator");
		this.animation = undefined;
	};
	Indicator.prototype =
		/** @lends ListHarvest.GUI */
		{
			set: function(percent) {
				if (this.animation) {
					this.animation.finish();
				}
				this.animation = this.element.find(".progress-bar").animate({
					"width": percent + "%"
				}, 100);
			},
		 stop: function () {
			 	 this.element.removeClass("progress-striped");
		 },
		 start: function () {
				 this.element.find(".progress-bar").removeClass("progress-bar-warning");
				 this.element.find(".progress-bar").addClass("progress-bar-default");
			 	 this.element.addClass("progress-striped");
		 },
		finished: function () {
				this.changeIndicator("success");
		},
		/**
		 * It changes the design of indicator for a warning situation
		 */
		warning: function () {
				this.changeIndicator("warning");
		},
		/**
		 * It changes the design of indicator for an error situation
		 */
		error: function () {
				this.changeIndicator("danger");
		},
		/**
		 * It changes the design of the indicator for a certain situation
		 * @param  {string} type The situation. It can be "danger" or "warning" or "success"
		 */
	 changeIndicator: function (type) {
			 this.element.find(".progress-bar").addClass("progress-bar-" + type);
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
