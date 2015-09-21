/* global $, wisply,window */
/**
 * @file Encapsulates the functionality for managing repositories
 * @author Cristian Sima
 */
/**
 * @namespace Repositories
 */
var Repositories = function() {
    'use strict';

    /**
     * Creates an empty history
     * @memberof Repositories
     * @class History
     * @classdesc It holds a history of events
     */
    var History = function History() {
        this.data = [];
    };

    History.prototype =
        /** @lends Repositories.History */
        {
            /**
             * It adds a log message
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
             * It adds an event to history. It assigns the datestamp
             * @param {string} content The content of the message
             * @param {string} type    The type of the message. It can be "LOG", "ERROR" or "WARN"
             */
            add: function(content, type) {
                var datetime;

                function getNiceDate() {
                    var currentdate = new Date(),
                        datetime = currentdate.getDate() + "." +
                        (currentdate.getMonth() + 1) + "." +
                        currentdate.getFullYear() + " at " +
                        currentdate.getHours() + ":" +
                        currentdate.getMinutes() + ":" +
                        currentdate.getSeconds();
                    return datetime;
                }

                datetime = getNiceDate();

                this.data.unshift({
                    date: datetime,
                    content: content,
                    type: type
                });
                if (wisply.repositories) {
                    wisply.repositories.page.update();
                }
            },
            /**
             * It returns the history in HTML format
             * @return [string] The history in HTML format
             */
            getHTML: function() {

                function getHeader() {
                    var header = "";
                    header = "<thead><tr><th class='text-center'>Date</th><th class='text-center'>Category</th><th class='text-center'>Content</th></tr></thead>";
                    return header;
                }

                function getBody(arrray) {
                    var result = "<tbody>",
                        i, currentEvent;


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
                                content = "Warn";
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
     * Starts the connection and inits the listeners
     * @memberof Repositories
     * @class Connection
     * @classdesc It represents a websocket connection
     * @param {function} processor A callback called when a message is received
     */
    var Connection = function Connection(processor) {

        this.value = new WebSocket("ws://" + host + "/admin/repositories/ws");

        this.processor = processor;
        this.initListeners();
    };

    Connection.prototype =
        /** @lends Repositories.Connection */
        {
            /**
             * It
             */
            initListeners: function() {
                this.value.onopen = function(evt) {
                    wisply.repositories.history.log("The websocket connection has been created");
                    $("#connectionStatus").html("<span class='text-success'>WebSocket connection established</span>");
                };
                this.value.onclose = function(evt) {
                    wisply.repositories.history.logError("The webscoket connection is closed");
                    $("#connectionStatus").html("<span class='text-danger'>No WebSocket connection</span>");
                    wisply.repositories.page.errorOcurred();
                };
                this.value.onmessage = this.processor;
                this.value.onerror = function(evt) {
                    wisply.repositories.history.logWarning("There was a an error with web scoket connection");
                };
            },
            /**
             * It sends a message
             * @param  {string} value The value of the message
             */
            sendMessage: function(name, value) {
                var msg = {
                    name: name,
                    value: value
                };
                this.value.send(JSON.stringify(msg));
            }
        };

    /**
     * Saves the stages
     * @memberof Repositories
     * @class StagerManager
     * @classdesc It encapsulets the functionality for the sources
     * @param [Manager] repositoriesManager The reference to the repositories manager
     */
    var StageManager = function StageManager(repositoriesManager) {
        this.repo = repositoriesManager;
        this.data = [{
            name: "Connect to server",
            id: 0,
            perform: function(stageManager) {
                var manager = stageManager.repo;
                if (window.WebSocket) {
                    manager.connection = new Connection(function(data) {
                        wisply.repositories.processMessage(data);
                    });
                } else {
                    alert("Your browser does not support WebSockets");
                    return;
                }
                setTimeout(function() {
                    stageManager.firedStageFinished();
                }, 400);
            }
        }, {
            name: "Test URL",
            id: 1,
            perform: function(stageManager) {
                var instance = stageManager;

                instance.repo.history.log("Indentifing the source");

                  instance.repo.connection.sendMessage("testURL", "something");

            },
            result: function(stageManager, content) {
                  console.log(content);
                if (content === "true") {
                    stageManager.repo.history.log("The URL is valid");
                    stageManager.firedStageFinished();
                } else {
                    stageManager.repo.history.log("The URL is not valid!");
                    stageManager.repo.page.errorOcurred();
                }
            }
        }, {
            name: "Identify Source",
            id: 2,
            perform: function(stageManager) {
                var instance = stageManager;

                instance.repo.history.log("Indentifing the source");

                instance.repo.connection.sendMessage("identify", "something");
            },
            result: function(content) {}

        }];
        this.current = "None";
        this.stage = {};
    };
    StageManager.prototype =
        /** @lends Repositories.StageManager */
        {
            start: function() {
                this.current = 0;
                this.performStage(0);
            },
            next: function() {
                this.current++;
                if (this.current >= this.data.length) {
                    this.firedEnd();
                } else {
                    this.repo.history.log("Starting stage " + (this.current + 1) + "...");
                    this.performStage(this.current);
                }
            },
            performStage: function(id) {
                var stage = this.data[id];
                this.stage = stage;
                stage.perform(this);
            },
            forceStop: function() {
                this.current = "None";
                this.stage.stop();
            },
            firedStageFinished: function() {
                this.repo.page.update();
                this.repo.history.log("Stage " + (this.current + 1) + " finished!");
                this.next();
            },
            firedEnd: function() {
                this.repo.page.update();
                this.repo.history.log("The process has been finished!");
            }
        };


    /**
     * The constructor activates the listeners
     * @memberof Repositories
     * @class Manager
     * @classdesc It encapsulets the functionality for the sources
     */
    var Manager = function Manager() {
        this.history = new History();
        this.history.log("The manager has started.");
        this.page = new Page();
        this.stageManager = new StageManager(this);
    };
    Manager.prototype =
        /** @lends Repositories.Manager */
        {
            /**
             * It activates the listeners
             */
            init: function() {
              var instance = this;
                this.page.update();
                setTimeout(function() {
                  instance.stageManager.start();
                }, 1000);
            },
            /**
             * It processes the messages received from the server
             */
            processMessage: function(evt) {
                var msg = JSON.parse(evt.data);

                function getContentMessage(content) {

                    if (content) {
                        return " with content [" + content + "]";
                    }

                    return ", which does not has content";
                }
                this.history.log("I received the socket [<b>" + msg.name + "</b>]" + getContentMessage(msg.content));
                this.chooseAction(msg.name, msg.content);
            },
            /**
             * It choose the action based on name
             * @param  {string} name    The name of the message
             * @param  {string} content The content of the message
             */
            chooseAction: function(name, content) {
                switch (name) {
                    case "identified":
                        this.stageManager.firedStageFinished();
                        break;
                    case "URLTested":
                        this.stageManager.stage.result(this.stageManager, content);
                        break;
                }
            }
        };

    /**
     *
     * @memberof Repositories
     * @class Page
     * @classdesc It manages the GUI of the page
     * @param {function} processor A callback called when a message is received
     */
    var Page = function Page(processor) {
        this.init();
        this.currentTab = "current";
    };

    Page.prototype =
        /** @lends Repositories.Page */
        {
            /**
             * It activates the listeners
             */
            init: function() {
                var instance = this;
                $("#historyButton").click(function() {
                    instance.showHistory();
                });
            },
            showHistory: function() {
                this.changeTab("history");
            },
            changeTab: function(newTab) {
                this.currentTab = newTab;
                this.update();
            },
            /**
             * It refreshes the history
             */
            refreshHistory: function() {
                $("#history").html(wisply.repositories.history.getHTML());
            },
            /**
             * It updates the current view
             */
            update: function() {
                switch (this.currentTab) {
                    case "current":
                        break;
                    case "history":
                        this.refreshHistory();
                        break;

                }
                this.updateStages();
                this.updateHistoryNumber();
            },
            /**
             * It updates the number of events in history
             */
            updateHistoryNumber: function() {
                var nr = wisply.repositories.history.data.length;
                $("#historyButton").find("#historyBadge").html(nr);
            },
            updateStages: function() {
                var manager = wisply.repositories.stageManager,
                    current = manager.current,
                    stages = manager.data,
                    stage, id, html = "",
                    item = "";
                html += "";
                for (id = 0; id < stages.length; id++) {
                    stage = stages[id];
                    if (id === current) {
                        item = '<li class="list-group-item active">' + stage.name + '</li>';
                    } else {
                        if (id < current) {
                            item = '<li class="list-group-item text-muted"><del>' + stage.name + '</del></li>';
                        } else {
                            item = '<li class="list-group-item">' + stage.name + '</li>';
                        }
                    }
                    html += item;
                }
                if (current === stages.length) {
                    html += '<div class="panel panel-primary">  <div class="panel-heading">    <h3 class="panel-title">Great!</h3></div>  <div class="panel-body">    The process is over.  </div></div>';
                    this.processFinished();
                }
                $("#stages").html(html);
                this.updateGeneralIndicator();
            },
            updateGeneralIndicator: function() {
                var manager = wisply.repositories.stageManager,
                    current = manager.current,
                    total = manager.data.length,
                    percent = 0;
                if (current === "NONE") {
                    percent = 0;
                } else {
                    percent = (current) / total;
                }
                percent = percent * 100;

                if (this.animateGeneralIndicator) {
                    this.animateGeneralIndicator.finish();
                }

                this.animateGeneralIndicator = $("#generalIndicator").find(".progress-bar").animate({
                    "width": percent + "%"
                }, 100);
            },
            processFinished: function() {
                var general = $("#generalIndicator");
                general.removeClass("progress-striped");
                general.find(".progress-bar").addClass("progress-bar-success");
            },
            errorOcurred: function() {
                var general = $("#generalIndicator");
                general.removeClass("progress-striped");
                general.find(".progress-bar").addClass("progress-bar-danger");
            }
        };


    return {
        Manager: Manager
    };
};
$(document).ready(function() {
    "use strict";
    var module = new Repositories();
    wisply.repositories = new module.Manager();
    wisply.repositories.init();
});
