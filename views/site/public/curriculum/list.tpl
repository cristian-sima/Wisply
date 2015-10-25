<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Curricula panel</li>
    </ul>
  </div>
  <div class="panel-body">
    <section>
      <div class="btn-group-sm">
        <a href="/admin/curriculum/add" class="btn btn-primary">
          <span class="glyphicon glyphicon-plus"></span> Add program of study</a>
        </div>
      </section>
      <br />
      <section>
        {{ if eq (len .programs) 0 }}
        There is no program of study.
        {{ else }}
        <div>
          <div class="row text-center">
            {{range $index, $program := .programs}}
            <div class="text-central col-xs-12 col-sm-6 col-md-3 col-ls-2" >
              <a href="/admin/curriculum/programs/{{ $program.GetID }}">
                <div style="height:100px;" class="thumbnail">
                  <div class="caption">
                    <h3>
                      {{ ($program.GetName) | html }}
                    </h3>
                  </div>
                </div>
              </a>
              </div>
            {{end }}
            </div>
          </div>

          <div id="harvest-history-container" class="modal fade">
            <div class="modal-dialog">
              <div class="modal-content">
                <div class="modal-header">
                  <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                  <h4 class="modal-title">History</h4>
                </div>
                <div class="modal-body" id="harvest-history-element">
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                </div>
              </div>
            </div>
          </div>
          {{ end }}
        </div>
      </div>
