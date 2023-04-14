# GoBro
![GoBro](https://user-images.githubusercontent.com/35332515/231014731-3e4afc93-690a-4cbb-b7dd-5cf1f0c0044c.png)
<h2>A Go browser. Render pages using new format ".ego" and standard ".css".</h2>
<p>Go is a beautiful programming language, so it deserves a browser written on it. But not <b><i>any</i></b> browser. Especially not the type that runs on <b><i>JavaScript</i></b>.</p>
<p>This <b>project</b> aims to create an entirely independent system of parsing and displaying documents.</p>
<p>For now, it lacks visuals, but already can create a tree of objects parsed from the file with custom extension "<b>.ego</b>".<br/>
It may not be the perfect solution for displaying pages at the moment, but it aims to replace it in the future.</p><br/>
<h3>The stage of the project</h3>
<ul>
<li>Parser _________________ 🗸</li>
<li>IPC ____________________ 🗸</li>
<li>Hashmaps ____________ 🗸</li>
<li>sPOT optimization ____ 🗸</li>
<li>Client API _____________ in progress</li>
<li>Standalone ___________ not started</li>
<li>GIO UI ________________ not started</li>
<li>HTTP _________________ not started</li>
<li>Public test ____________ not started</li>
</ul>
<h3>Now to the rules:</h2>
<p>Modern standard requires high levels of readability, adaptability, functionality, and must provide total freedom for designers and maintainers.</p>
<p>This is why we try to get rid of the complexity of parsing the DOM, so no <b>div</b><i>s</i>, <b>p</b><i>s</i>, <b>h</b><i>s</i>, and other madness (<i>just joking about madness</i>).</p>
<p>We will provide only 4 possible <b>tag</b><i>s</i>:</p>
<ul>
<li><b>&lt;s&gt;</b> - links a custom style to the objects,</li>
<li><b>&lt;x&gt;</b> - links a custom executable to the page,</li>
<li><b>###</b> - just a comment line,</li>
<li><b>&lt;e&gt;</b> - prototype of the element.</li>
</ul>
<p>Yes, you saw it right! One <b>&lt;e&gt;</b> to rule them all!</p>
<p>It gets even better:</p>
<ul>
<li><b>###</b> comments are one-liners, so everything on the line is omitted during parsing,</li>
<li><b>&lt;s&gt;</b> and <b>&lt;x&gt;</b> have only one field - <b>ref</b>, so it's declared as follows: <b>&lt;s ref="style.css"&gt;</b> or <b>&lt;x ref="executable.go"&gt;</b>. Every style and executable is linked and executed one by one in descending order, with the top links or executes first,</li>
<li><b>&lt;e&gt;</b> can have <b>&lt;e id="id" class="class" ref="reference_to_something"&gt;</b> properties and must be closed with <b>&lt;/&gt;</b></li>
</ul>
<p>There is the best part: <b>&lt;e&gt;</b> can have only one or none properties. The <b>id</b> and <b>class</b> are used to link an object to the corresponding style provided with <b>style.css</b>. If no <b>id</b> and <b>class</b> is provided, linker will give it <b>std</b> property, that is the plain text.</p>
<i><p>In later versions, <b>context</b> will be added, which will be key in providing context for the linker. For example, <b>context="text"</b> will treat object like a plain text, <b>context="button"</b> as button, <b>context="field"</b> as input field, etc. It will also support "user" definition that will tell linker to treat objects with <b>id</b> or <b>class</b> differently. Be aware, that with "user" and no <b>id</b> or <b>class</b> provided, it will still be treated as plain text!</p></i>
