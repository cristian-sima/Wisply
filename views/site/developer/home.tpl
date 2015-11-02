<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <aside class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li class="active">API &amp; Developers</li>
          </ul>
        </aside>
        <section class="panel-body">
          <div>
            <div class="row">
              <div class="col-lg-6 col-md-6 col-sm-6">
                <h1>Developers &amp; Research</h1>
                <img class="img-responsive" src="/static/img/developer/cloud.png" alt="API &amp; Developers" />
              </div>
              <div class="col-lg-6 col-md-6 col-sm-6">
                <blockquote>
                  <h2>
                    <span class="text-muted">
                      <span class="glyphicon glyphicon-align-left"></span>
                  </span>
                   Free to
                    <span class="text-muted">re</span>-use
                  </h2>
                  If you want to understand how Wisply is spinning the wheels, you can download the source code from our <a href="https://github.com/cristian-sima/Wisply" target="_blank">GitHub page</a>.
                </blockquote>
                <blockquote>
                  <h2>
                    <span class="text-muted">
                      <span class="glyphicon glyphicon-stats"></span>
                  </span>
                  Open data
                  </h2>
                  Wisply is making available all possible data in <a class="scroll" href="#open-data">open data</a> format.
                </blockquote>
                <blockquote>
                  <h2><span class="text-muted"><span class="glyphicon glyphicon-dashboard"></span></span> Research tools</h2>
                  We provide a <a class="scroll" href="#tools">list of tools</a> which Wisply is using and you may find useful for your research.
                </blockquote>
              </div>
            </div>
          </div>
          <hr />
          <br />
        </section>
        <div>
        <section id="open-data">
          <div class="panel panel-default">
            <div class="panel-heading">
              <h3 class="panel-title">Download data sets and tables</h3>
            </div>
            <div class="panel-body">
              <br />
              In case you want to use the data, you can download the entire tables which Wisply is using. <a href="/developer/data/table" class="btn btn-primary btn-xs">See the list of tables</a>
              <br />
            </div>
          </div>
        </section>
        <section>
          <div class="panel panel-default">
            <div class="panel-heading">
              <h3 id="http" class="panel-title">API - HTTP requests</h3>
            </div>
            <div class="panel-body">
              <br />
              You can integrate data from Wisply with your application. Wisply makes available a public API which you can use.
              <br />
              <br />
              <ul class="list-group">
                <li class="list-group-item">
                  <div>
                    <h4>List the resources from a repository</h4>
                    Type <span class="label label-info">GET</span>
                    <br />
                    <pre>developer/api/repository/resources/{<strong>repositoryID</strong>}/get/{<strong>startResource</strong>}/{<strong>resourceNumber</strong>}?collection={<strong>collectionID</strong>}&amp;format={<strong>format</strong>}</pre>
                    Response format: <span class="label label-success">JSON</span> or  <span class="label label-success">HTML</span>
                    <br />
                    Where:
                    <br />
                    <ul>
                      <li><strong>repositoryID</strong> is the ID of the repository</li>
                      <li><strong>startResource</strong> is the wisply ID of the first resource</li>
                      <li><strong>resourceNumber</strong> the number of resources. Please note  that Wisply limits the number to 100. For more, please download the table</li>
                      <li><strong>collectionId</strong> is the id of the collection within the repository</li>
                      <li><strong>format</strong> may be html or json</li>
                    </ul>
                  </div>
                </li>
                <li class="list-group-item">
                  <div>
                    <h4>Search for items</h4>
                    Type <span class="label label-info">GET</span>
                    <br />
                    <pre>developer/api/search/anything/{<strong>query</strong>}</pre>
                    Response format: <span class="label label-success">JSON</span>
                    <br />
                    <blockquote>
                      This request gets (in this order): <br />
                      <ul>
                        <li>curricula</li>
                        <li>institutions</li>
                        <li>repositories</li>
                        <li>collections</li>
                        <li>resources (in case the query is at least 5 characters)</li>
                      </ul>
                      Wisply returns at most 5 items of each type.
                    </blockquote>
                    <br />
                    Where: <br />
                    <ul>
                      <li><strong>query</strong> is the string which you want to search. It must be less then 100 characters.</li>
                    </ul>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </section>
        <section id="tools">
          <div class="panel panel-default">
            <div class="panel-heading">
              <h3 class="panel-title">Tools</h3>
            </div>
            <div class="panel-body">
              <br >
              <ol>
                <li class="h4"><h4>Digester</h4>
                  <div class="well h5">
                    <span class="text-muted glyphicon glyphicon glyphicon-tasks"></span> Digester is a simple tool which takes a text and produces a list of occurences for its words.
                    <br />
                    <br />
                    <a class="btn btn-primary" href="/developer/tools/digester">Access tool</a>
                  </div>
                </li>
              </ol>
              <br />
            </div>
          </div>
        </section>
      </div>
      </div>
    </div>
  </div>
</div>
