<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Curriculum panel</li>
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
      <h4>Programs of study</h4>
    <section>
      {{ if eq (len .programs) 0 }}
      There is no program of study.
      {{ else }}
      <div class="table-responsive">
        <table class="table table-striped table-hover " id="repositories-list">
          <thead>
            <tr>
              <th>Name</th>
            </tr>
          </thead>
          <tbody>
            {{range $index, $program := .programs}}
            <tr>
              <td><a class="btn btn-default" href="/admin/curriculum/programs/{{ $program.GetID }}">{{ ($program.GetName) | html }}</a></td>
            </tr>
            {{end }}
          </tbody>
        </table>
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

      </div>
      {{ end }}
    </div>
  </div>
