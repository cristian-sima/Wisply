<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <aside class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/api">API &amp; Developers</a></li>
          </ul>
        </aside>
        <section class="panel-body">
          <div class="row">
            <div class="col-lg-6 col-md-6 col-sm-6">
              <img class="img-responsive" src="/static/img/api/cloud.png" alt="API Image" />
            </div>
            <div class="col-lg-6 col-md-6 col-sm-6">
              <br />
              <blockquote>
                <h3>Open data</h3>
                Wisply is making available all possible data in <a target="_blank" href="http://theodi.org/">open data</a> format.
              </blockquote>
              <br />
              <blockquote>
                <h3>Open source</h3>
                If you want to understand how Wisply is spinning the wheels, you may want to go to our <a href="https://github.com/cristian-sima/Wisply" target="_blank">GitHub page</a>.
              </blockquote>
            </div>
          </div>
          <hr />
          <br />
        </section>
          <section>
            <div class="panel panel-default">
              <div class="panel-heading">
                <h3 class="panel-title">Download data sets and tables</h3>
              </div>
              <div class="panel-body">
                <br />
                You can download the entire tables which Wisply is using. The list is available <a href="/api/table/list" class="btn btn-primary btn-xs">here</a>
                <br />
              </div>
            </div>
          </section>
          <section>
            <div class="panel panel-default">
              <div class="panel-heading">
                <h3 id="http" class="panel-title">HTTP requests</h3>
              </div>
              <div class="panel-body">
                <br />
                You can integrate data from Wisply with your application. Wisply makes available a public API which you can use.
                <br />
                <br />
                <ul class="list-group">
                  <li class="list-group-item">
                    <div>
                      <h5>List the resources from a repository</h5>
                      Type <span class="label label-info">GET</span>
                      <br />
                      <pre>api/repository/resources/{<strong>repositoryID</strong>}/get/{<strong>startResource</strong>}/{<strong>resourceNumber</strong>}?collection={<strong>collectionID</strong>}&amp;format={<strong>format</strong>}</pre>
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
                      <h5>Search for items</h5>
                      Type <span class="label label-info">GET</span>
                      <br />
                      <pre>api/search/anything/{<strong>query</strong>}</pre>
                      <br />
                      <blockquote>
                        This request searches for (in this order): <br />
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
        </div>
      </div>
    </div>
  </div>
