package cameo

// No need to depend on external assets since
// because we there are only few dependencies
// so we can have them as static variables.

const Favicon = `
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
`

const FormTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<link rel="icon" href="data:image/png;base64,{{.Favicon}}"/>
<title>Cameo | {{.Title}}</title>
</head>
<body>

<form method="post">
<p style="text-align: center; margin-bottom: 40px">
<img src="data:image/png;base64,{{.Favicon}}" alt="{{.Title}}"/>
</p>
{{if .Note}}
<span class="note note-{{if .IsError}}error{{else}}info{{end}}">
{{.Note}}
</span>
{{end}}
<label for="subject">Subject</label>
<input type="text" name="subject" id="subject" value="{{.Message.Subject}}"/>

<label for="body">Your message</label>
<textarea name="body" id="body" rows="8">{{.Message.Body}}</textarea>
<button type="submit" class="button button-black float-right">Send</button>
</form>

</body>
<style type="text/css">
*,*:after,*:before{box-sizing:inherit}html{box-sizing:border-box;font-size:62.5%}body{color:#606c76;font-family:'Roboto', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;font-size:1.6em;font-weight:300;letter-spacing:.01em;line-height:1.6}blockquote{border-left:0.3rem solid #d1d1d1;margin-left:0;margin-right:0;padding:1rem 1.5rem}blockquote *:last-child{margin-bottom:0}.button,button,input[type='button'],input[type='reset'],input[type='submit']{background-color:#9b4dca;border:0.1rem solid #9b4dca;border-radius:.4rem;color:#fff;cursor:pointer;display:inline-block;font-size:1.1rem;font-weight:700;height:3.8rem;letter-spacing:.1rem;line-height:3.8rem;padding:0 3.0rem;text-align:center;text-decoration:none;text-transform:uppercase;white-space:nowrap}.button:focus,.button:hover,button:focus,button:hover,input[type='button']:focus,input[type='button']:hover,input[type='reset']:focus,input[type='reset']:hover,input[type='submit']:focus,input[type='submit']:hover{background-color:#606c76;border-color:#606c76;color:#fff;outline:0}.button[disabled],button[disabled],input[type='button'][disabled],input[type='reset'][disabled],input[type='submit'][disabled]{cursor:default;opacity:.5}.button[disabled]:focus,.button[disabled]:hover,button[disabled]:focus,button[disabled]:hover,input[type='button'][disabled]:focus,input[type='button'][disabled]:hover,input[type='reset'][disabled]:focus,input[type='reset'][disabled]:hover,input[type='submit'][disabled]:focus,input[type='submit'][disabled]:hover{background-color:#9b4dca;border-color:#9b4dca}.button.button-outline,button.button-outline,input[type='button'].button-outline,input[type='reset'].button-outline,input[type='submit'].button-outline{background-color:transparent;color:#9b4dca}.button.button-outline:focus,.button.button-outline:hover,button.button-outline:focus,button.button-outline:hover,input[type='button'].button-outline:focus,input[type='button'].button-outline:hover,input[type='reset'].button-outline:focus,input[type='reset'].button-outline:hover,input[type='submit'].button-outline:focus,input[type='submit'].button-outline:hover{background-color:transparent;border-color:#606c76;color:#606c76}.button.button-outline[disabled]:focus,.button.button-outline[disabled]:hover,button.button-outline[disabled]:focus,button.button-outline[disabled]:hover,input[type='button'].button-outline[disabled]:focus,input[type='button'].button-outline[disabled]:hover,input[type='reset'].button-outline[disabled]:focus,input[type='reset'].button-outline[disabled]:hover,input[type='submit'].button-outline[disabled]:focus,input[type='submit'].button-outline[disabled]:hover{border-color:inherit;color:#9b4dca}.button.button-clear,button.button-clear,input[type='button'].button-clear,input[type='reset'].button-clear,input[type='submit'].button-clear{background-color:transparent;border-color:transparent;color:#9b4dca}.button.button-clear:focus,.button.button-clear:hover,button.button-clear:focus,button.button-clear:hover,input[type='button'].button-clear:focus,input[type='button'].button-clear:hover,input[type='reset'].button-clear:focus,input[type='reset'].button-clear:hover,input[type='submit'].button-clear:focus,input[type='submit'].button-clear:hover{background-color:transparent;border-color:transparent;color:#606c76}.button.button-clear[disabled]:focus,.button.button-clear[disabled]:hover,button.button-clear[disabled]:focus,button.button-clear[disabled]:hover,input[type='button'].button-clear[disabled]:focus,input[type='button'].button-clear[disabled]:hover,input[type='reset'].button-clear[disabled]:focus,input[type='reset'].button-clear[disabled]:hover,input[type='submit'].button-clear[disabled]:focus,input[type='submit'].button-clear[disabled]:hover{color:#9b4dca}code{background:#f4f5f6;border-radius:.4rem;font-size:86%;margin:0 .2rem;padding:.2rem .5rem;white-space:nowrap}pre{background:#f4f5f6;border-left:0.3rem solid #9b4dca;overflow-y:hidden}pre>code{border-radius:0;display:block;padding:1rem 1.5rem;white-space:pre}hr{border:0;border-top:0.1rem solid #f4f5f6;margin:3.0rem 0}input[type='color'],input[type='date'],input[type='datetime'],input[type='datetime-local'],input[type='email'],input[type='month'],input[type='number'],input[type='password'],input[type='search'],input[type='tel'],input[type='text'],input[type='url'],input[type='week'],input:not([type]),textarea,select{-webkit-appearance:none;background-color:transparent;border:0.1rem solid #d1d1d1;border-radius:.4rem;box-shadow:none;box-sizing:inherit;height:3.8rem;padding:.6rem 1.0rem .7rem;width:100%}input[type='color']:focus,input[type='date']:focus,input[type='datetime']:focus,input[type='datetime-local']:focus,input[type='email']:focus,input[type='month']:focus,input[type='number']:focus,input[type='password']:focus,input[type='search']:focus,input[type='tel']:focus,input[type='text']:focus,input[type='url']:focus,input[type='week']:focus,input:not([type]):focus,textarea:focus,select:focus{border-color:#9b4dca;outline:0}select{background:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 8" width="30"><path fill="%23d1d1d1" d="M0,0l6,8l6-8"/></svg>') center right no-repeat;padding-right:3.0rem}select:focus{background-image:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 8" width="30"><path fill="%239b4dca" d="M0,0l6,8l6-8"/></svg>')}select[multiple]{background:none;height:auto}textarea{min-height:6.5rem}label,legend{display:block;font-size:1.6rem;font-weight:700;margin-bottom:.5rem}fieldset{border-width:0;padding:0}input[type='checkbox'],input[type='radio']{display:inline}.label-inline{display:inline-block;font-weight:normal;margin-left:.5rem}.container{margin:0 auto;max-width:112.0rem;padding:0 2.0rem;position:relative;width:100%}.row{display:flex;flex-direction:column;padding:0;width:100%}.row.row-no-padding{padding:0}.row.row-no-padding>.column{padding:0}.row.row-wrap{flex-wrap:wrap}.row.row-top{align-items:flex-start}.row.row-bottom{align-items:flex-end}.row.row-center{align-items:center}.row.row-stretch{align-items:stretch}.row.row-baseline{align-items:baseline}.row .column{display:block;flex:1 1 auto;margin-left:0;max-width:100%;width:100%}.row .column.column-offset-10{margin-left:10%}.row .column.column-offset-20{margin-left:20%}.row .column.column-offset-25{margin-left:25%}.row .column.column-offset-33,.row .column.column-offset-34{margin-left:33.3333%}.row .column.column-offset-40{margin-left:40%}.row .column.column-offset-50{margin-left:50%}.row .column.column-offset-60{margin-left:60%}.row .column.column-offset-66,.row .column.column-offset-67{margin-left:66.6666%}.row .column.column-offset-75{margin-left:75%}.row .column.column-offset-80{margin-left:80%}.row .column.column-offset-90{margin-left:90%}.row .column.column-10{flex:0 0 10%;max-width:10%}.row .column.column-20{flex:0 0 20%;max-width:20%}.row .column.column-25{flex:0 0 25%;max-width:25%}.row .column.column-33,.row .column.column-34{flex:0 0 33.3333%;max-width:33.3333%}.row .column.column-40{flex:0 0 40%;max-width:40%}.row .column.column-50{flex:0 0 50%;max-width:50%}.row .column.column-60{flex:0 0 60%;max-width:60%}.row .column.column-66,.row .column.column-67{flex:0 0 66.6666%;max-width:66.6666%}.row .column.column-75{flex:0 0 75%;max-width:75%}.row .column.column-80{flex:0 0 80%;max-width:80%}.row .column.column-90{flex:0 0 90%;max-width:90%}.row .column .column-top{align-self:flex-start}.row .column .column-bottom{align-self:flex-end}.row .column .column-center{align-self:center}@media (min-width: 40rem){.row{flex-direction:row;margin-left:-1.0rem;width:calc(100% + 2.0rem)}.row .column{margin-bottom:inherit;padding:0 1.0rem}}a{color:#9b4dca;text-decoration:none}a:focus,a:hover{color:#606c76}dl,ol,ul{list-style:none;margin-top:0;padding-left:0}dl dl,dl ol,dl ul,ol dl,ol ol,ol ul,ul dl,ul ol,ul ul{font-size:90%;margin:1.5rem 0 1.5rem 3.0rem}ol{list-style:decimal inside}ul{list-style:circle inside}.button,button,dd,dt,li{margin-bottom:1.0rem}fieldset,input,select,textarea{margin-bottom:1.5rem}blockquote,dl,figure,form,ol,p,pre,table,ul{margin-bottom:2.5rem}table{border-spacing:0;display:block;overflow-x:auto;text-align:left;width:100%}td,th{border-bottom:0.1rem solid #e1e1e1;padding:1.2rem 1.5rem}td:first-child,th:first-child{padding-left:0}td:last-child,th:last-child{padding-right:0}@media (min-width: 40rem){table{display:table;overflow-x:initial}}b,strong{font-weight:bold}p{margin-top:0}h1,h2,h3,h4,h5,h6{font-weight:300;letter-spacing:-.1rem;margin-bottom:2.0rem;margin-top:0}h1{font-size:4.6rem;line-height:1.2}h2{font-size:3.6rem;line-height:1.25}h3{font-size:2.8rem;line-height:1.3}h4{font-size:2.2rem;letter-spacing:-.08rem;line-height:1.35}h5{font-size:1.8rem;letter-spacing:-.05rem;line-height:1.5}h6{font-size:1.6rem;letter-spacing:0;line-height:1.4}img{max-width:100%}.clearfix:after{clear:both;content:' ';display:table}.float-left{float:left}.float-right{float:right}
body {
width: 60%;
margin: 80px auto;
}

textarea {
min-height: 180px;
}

form {
padding: 10px;
}

label {
font-weight: 400;
}

.button-black {
background-color: black;
border-color: black;
}

.note {
display: inline-block;
border-radius: 1px;
margin: 20px 0;
padding: 5px 8px;
}

.note-error {
background: #ca2a2a;
color: #fbcfcf;
}

.note-info {
background: #2a7fca;
color: #cfedfb;
}
</style>
</html>
<!-- Version: {{.Version}} -->
`
