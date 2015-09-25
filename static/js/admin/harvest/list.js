/* global $, wisply,window, data, server*/
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
              console.log("Hi. It works :)");
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
                    this.GUI.changeAllStatus(message.Value);
                    break;
                case "RepositoryChangedStatus":
                    this.GUI.changeStatus({
                        id: message.Repository,
                        status: message.Content.NewStatus
                    });
                    break;
                }
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
                for (index = 0; index <= repositories.length; i++) {
                    repository = repositories[index];
                    this.changeStatus(repository);
                }
            },
            changeStatus: function (repository) {
                var htmlID = this.getHTMLID(repository.id);
                this.list.find(htmlID).html(repository.status);
            },
            getHTMLID: function (id) {
                return "rep-status-" + id;
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
    repository = new Repositories();
    decision = new list.DecisionManager();
    stages = list.Stages;
    stage = new harvest.StageManager(stages);
    manager = new harvest.Manager(stage, decision);
    manager.start();
});
