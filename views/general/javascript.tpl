{{ define "javascript" }}
<script src="/static/paper/jquery-1.10.2.min.js"></script>
<script src="/static/js/bootbox.min.js"></script>
<script src="/static/paper/bootstrap.min.js"></script>
<script src="/static/paper/bootswatch.js"></script>
<script type="text/javascript">
  /* <![CDATA[ */
  (function (){try{var s,a,i,j,r,c,l=document.getElementsByTagName("a"),t=document.createElement("textarea");for(i=0;l.length-i;i++){try{a=l[i].getAttribute("href");if(a&&a.indexOf("/cdn-cgi/l/email-protection") > -1  && (a.length > 28)){s='';j=27+ 1 + a.indexOf("/cdn-cgi/l/email-protection");if (a.length > j) {r=parseInt(a.substr(j,2),16);for(j+=2;a.length>j&&a.substr(j,1)!='X';j+=2){c=parseInt(a.substr(j,2),16)^r;s+=String.fromCharCode(c);}j+=1;s+=a.substr(j,a.length-j);}t.innerHTML=s.replace(/</g,"&lt;").replace(/>/g,"&gt;");l[i].setAttribute("href","mailto:"+t.value);}}catch(e){}}}catch(e){}}
  )();
  /* ]]> */
</script>
<script src="/static/js/base64_decode.js"></script>
<script src="/static/js/jquery.cookie.js"></script>

<script src="/static/js/general/init.js"></script>

{{ end }}
