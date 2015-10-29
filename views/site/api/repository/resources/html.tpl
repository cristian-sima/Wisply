<table class="table" id="listOfRecords">
  <tbody>
    {{range $index, $record := .records}}
    <tr>
      <td>
          <a class="resource" href="{{ $record.GetWisplyURL }}">
              <h4>
            {{range $index, $title := $record.Keys.Get "title" }}
            {{ $title }}
            {{ end }}
          </h4>
          </a>
        {{range $index, $description := $record.Keys.Get "identifier" }}
        <!-- <span class="label label-info">{{ $description }} </span><br /> -->
        {{ end }}

        {{range $index, $description := $record.Keys.Get "description" }}
        {{ $description }}
        {{ end }}
        <div class="formats">
          {{range $format, $number := $record.Keys.ProcessFormats }}
          <span class="label label-default">{{ $format}}{{ if ne $number 1}}<span class=" text-almost-invisible"> <span class="small">&times</span> </span> <span class="">{{ $number }}</span>
          {{ end }}</span> &nbsp;
          {{ end }}
        </div>
        <div class="creators">
          <span class="text-muted">
            {{range $index, $creator := $record.Keys.Get "creator" }}
            {{ $creator }}
            {{ end }}
          </span>
        </div>
      </td>
    </tr>
    {{ end }}
  </tbody>
</table>
