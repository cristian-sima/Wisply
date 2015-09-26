/* global $, Harvest, wisply*/
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Harvest
 */
var HarvestList = function () {
    'use strict';
    var Stages = [{
          id: 0,
          perform: function(manager) {
              this.manager = manager;
              this.manager.sendMessage("getAllRepositoriesStatus");
          }
    }];

    var DecisionManager = function DecisionManager() {
        this.GUI = new GUI();
    };
    DecisionManager.prototype =
        /** @lends Harvest.Manager */
        {
            decide: function (message) {
                switch (message.Name) {
                case "ListRepositoriesStatus":
                    this.changeAllStatus(message);
                    break;
                case "RepositoryChangedStatus":
                    this.changeSingleStatus(message);
                    break;
                }
            },
            changeAllStatus: function (message) {
                this.GUI.changeAllStatus(message.Value);
            },
            changeSingleStatus: function (message) {
                this.GUI.changeStatus({
                    id: message.Repository,
                    status: message.Value.NewStatus
                });
                this.GUI.activateActionListeners();
            }
        };
    var GUI = function GUI() {
        this.list = $("#repositories-list");
    };
    GUI.prototype =
        /** @lends ListHarvest.GUI */
        {
            changeAllStatus: function (repositories) {
                var repository, index;
                for (index = 0; index <= repositories.length; index++) {
                    repository = repositories[index];
                    this.changeStatus(repository);
                }
                this.activateActionListeners();
            },
            changeStatus: function (repository) {
                var htmlID = this.getHTMLID(repository.id),
                  htmlSpan = this.getStatusColor(repository.status),
                  action = this.getAction(repository);
                this.list.find(htmlID).html(htmlSpan + action);
            },
            activateActionListeners: function () {
                wisply.repositoriesModule.GUI.activateActionListeners();
            },
            getStatusColor: function (status) {
              return wisply.repositoriesModule.GUI.getStatusColor(status);
            },
            getAction: function(repository) {
              var action = "";
                switch(repository.status) {
                  case "unverified":
                    action = "<a href=''> <span data-toggle='tooltip' data-ID='" + repository.id + "' data-placement='top' title='' data-original-title='Start now!' class='repositories-init-harvest glyphicon glyphicon-sort-by-attributes hover' ></span></a>";
                  break;
                  case "verification-failed":
                  action = "<a href=''> <span data-toggle='tooltip' data-ID='" + repository.id + "' data-placement='top' title='' data-original-title='Try again' class='repositories-init-harvest glyphicon glyphicon glyphicon-refresh hover' ></span></a>";

                }
                return action;
            },
            getHTMLID: function (id) {
                return "#rep-status-" + id;
            }
        };
    return {
        DecisionManager: DecisionManager,
        Stages: Stages
    };
};
$(document).ready(function () {
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
