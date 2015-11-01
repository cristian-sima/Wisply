<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li class="active">Education</li>
            <li class="active">Programs of study</li>
          </ul>
        </div>
        <section class="panel-body">
            <h1>Programs of study</h1>
            <br />
            <div>
              <div class="row text-center">
                {{range $index, $program := .programs}}
                <div class="text-central col-xs-12 col-sm-6 col-md-3 col-ls-2" >
                  <a href="/education/programs/{{ $program.GetID }}">
                    <article style="height:150px;" class="thumbnail">
                      <span class="glyphicon glyphicon-bookmark big-icon"></span>
                      <div class="caption">
                        <h2>
                          {{ ($program.GetName) | html }}
                        </h2>
                      </div>
                    </article>
                  </a>
                </div>
                {{end }}
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
  <div>
    <style scoped>
    .big-icon {
      font-size: 56px;
    }
    </style>
  </div>
