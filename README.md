# GoBro
![GoBro](https://user-images.githubusercontent.com/35332515/230258554-7fe9de2c-55c9-4142-8df8-15aea8eff420.png)
<h2>A Go browser. Render pages using new format ".ego" and standard ".css".</h2>
<p>Go is beautiful programming language, so it deserves a browser written on it. But not <b><i>any</i></b> browser. Especially not that type that runs on <b><i>JavaScript</i></b>.</p>
<p>This <b>project</b> aims to create an entirely independent system of parsing and displaying documents.</p><br/>
<p>For now, it lacks visuals, but already can create a tree of objects parsed from the file with custom extension "<b>.ego</b>".<br/>
It may be not the perfect solution for displaying pages <i>(for now)</i>, but it will try to replace it in future.</p><br/>
<h3>Now to the rules:</h2>
<p>Modern standard requires high levels of readability, adaptability, functionality, and must provide total freedom for designers and maintainers.</p>
<p>This why we try to get rid of DOM's parsing complexity, so no <b>div</b><i>s</i>, <b>p</b><i>s</i>, <b>h</b><i>s</i>, and other madness (<i>just joking about madness</i>).</p>
<p>We will provide 4 and only possible <b>tag</b><i>s</i>:</p>
<ul>
<li><b>&lt;s&gt;</b> - links a custom style to the objects,</li>
<li><b>&lt;x&gt;</b> - links a custom executable to the page,</li>
<li><b>###</b> - just a comment line,</li>
<li><b>&lt;e&gt;</b> - prototype of the element.</li>
</ul>
<p>Yes, you saw it right! One <b>&lt;e&gt;</b> to rule them all!</p>
<p>It gets even better:</p>
<ul>
<li><b>###</b> comments are one-liners, so everything on the line is ommited when parsed,</li>
<li><b>&lt;s&gt;</b> and <b>&lt;x&gt;</b> have only one field - <b>ref</b>, so it's declared as follows: <b>&lt;s ref="style.css"&gt;</b> or <b>&lt;x ref="executable.go"&gt;</b>. Every style and executable is linked and executed one-by-one in descending order. So the top links and executes first,</li>
<li><b>&lt;e&gt;</b> can have <b>&lt;e id="id" class="class" ref="reference_to_something"&gt;</b> properties and must be closed with <b>&lt;/&gt;</b></li>
</ul>
<p>There is the best part: <b>&lt;e&gt;</b> can have only one or none properties. The <b>id</b> and <b>class</b> are used to link an object to the corresponding style provided with <b>style.css</b>. If no <b>id</b> and <b>class</b> were provided, linker will give it <b>std</b> property, that is the plain text.</p>
<i><p>In later versions there will be <b>context</b> added, that will be the key to provide context for the linker. For example, <b>context="text"</b> will treat object like a plain text, <b>context="button"</b> as button, <b>context="field"</b> as input field, etc. It will also support "user" definition, that will tell linker to treat objects with <b>id</b> or <b>class</b> totally different. Be aware, that with "user" and no <b>id</b> or <b>class</b> provided, it will still be treated as plain text!</p></i>
