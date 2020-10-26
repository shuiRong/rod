package js

var Element = &Function{
	Name:         "element",
	Definition:   `function(...e){const t=functions.selectable(this);for(const n of e){const e=t.querySelector(n);if(e)return e}return null}`,
	Dependencies: []*Function{Selectable},
}

var Elements = &Function{
	Name:         "elements",
	Definition:   `function(e){return functions.selectable(this).querySelectorAll(e)}`,
	Dependencies: []*Function{Selectable},
}

var ElementX = &Function{
	Name:         "elementX",
	Definition:   `function(...e){const t=functions.selectable(this);for(const n of e){const e=document.evaluate(n,t,null,XPathResult.FIRST_ORDERED_NODE_TYPE).singleNodeValue;if(e)return e}return null}`,
	Dependencies: []*Function{Selectable},
}

var ElementsX = &Function{
	Name:         "elementsX",
	Definition:   `function(e){const t=functions.selectable(this),n=document.evaluate(e,t,null,XPathResult.ORDERED_NODE_ITERATOR_TYPE),i=[];let o;for(;o=n.iterateNext();)i.push(o);return i}`,
	Dependencies: []*Function{Selectable},
}

var ElementR = &Function{
	Name:         "elementR",
	Definition:   `function(...e){for(let t=0;t<e.length-1;t+=2){const n=e[t],i=e[t+1],o=new RegExp(i),s=Array.from((this.document||this).querySelectorAll(n)).find(e=>o.test(functions.text.call(e)));if(s)return s}return null}`,
	Dependencies: []*Function{Text},
}

var Parents = &Function{
	Name:         "parents",
	Definition:   `function(e){let t=this.parentElement;const n=[];for(;t;)t.matches(e)&&n.push(t),t=t.parentElement;return n}`,
	Dependencies: []*Function{},
}

var ContainsElement = &Function{
	Name:         "containsElement",
	Definition:   `function(e){for(var t=e;null!=t;){if(t===this)return!0;t=t.parentElement}return!1}`,
	Dependencies: []*Function{},
}

var InitMouseTracer = &Function{
	Name:         "initMouseTracer",
	Definition:   `async function(e,t){if(await functions.waitLoad(),document.getElementById(e))return;const n=document.createElement("div");n.innerHTML=t;const i=n.lastChild;i.id=e,i.style="position: absolute; z-index: 2147483647; width: 17px; pointer-events: none;",i.removeAttribute("width"),i.removeAttribute("height"),document.body.appendChild(i)}`,
	Dependencies: []*Function{WaitLoad},
}

var UpdateMouseTracer = &Function{
	Name:         "updateMouseTracer",
	Definition:   `function(e,t,n){const i=document.getElementById(e);return!!i&&(i.style.left=t-2+"px",i.style.top=n-3+"px",!0)}`,
	Dependencies: []*Function{},
}

var Rect = &Function{
	Name:         "rect",
	Definition:   `function(){const e=functions.tag(this).getBoundingClientRect();return{x:e.x,y:e.y,width:e.width,height:e.height}}`,
	Dependencies: []*Function{Tag},
}

var Overlay = &Function{
	Name:         "overlay",
	Definition:   `async function(e,t,n,i,o,s){await functions.waitLoad();const r=document.createElement("div");if(r.id=e,r.style=` + "`" + `position: fixed; z-index:2147483647; border: 2px dashed red;\n        border-radius: 3px; box-shadow: #5f3232 0 0 3px; pointer-events: none;\n        box-sizing: border-box;\n        left: ${t}px;\n        top: ${n}px;\n        height: ${o}px;\n        width: ${i}px;` + "`" + `,i*o==0&&(r.style.border="none"),!s)return void document.body.appendChild(r);const l=document.createElement("div");l.style=` + "`" + `position: absolute; color: #cc26d6; font-size: 12px; background: #ffffffeb;\n        box-shadow: #333 0 0 3px; padding: 2px 5px; border-radius: 3px; white-space: nowrap;\n        top: ${o}px;` + "`" + `,l.innerHTML=s,r.appendChild(l),document.body.appendChild(r),window.innerHeight<l.offsetHeight+n+o&&(l.style.top=-l.offsetHeight-2+"px"),window.innerWidth<l.offsetWidth+t&&(l.style.left=window.innerWidth-l.offsetWidth-t+"px")}`,
	Dependencies: []*Function{WaitLoad},
}

var ElementOverlay = &Function{
	Name:         "elementOverlay",
	Definition:   `async function(e,t){const n=functions.tag(this);let i=n.getBoundingClientRect();await functions.overlay(e,i.left,i.top,i.width,i.height,t);const o=()=>{const t=document.getElementById(e);if(null===t)return;const s=n.getBoundingClientRect();i.left!==s.left||i.top!==s.top||i.width!==s.width||i.height!==s.height?(t.style.left=s.left+"px",t.style.top=s.top+"px",t.style.width=s.width+"px",t.style.height=s.height+"px",i=s,setTimeout(o,100)):setTimeout(o,100)};setTimeout(o,100)}`,
	Dependencies: []*Function{Tag, Overlay},
}

var RemoveOverlay = &Function{
	Name:         "removeOverlay",
	Definition:   `function(e){const t=document.getElementById(e);t&&t.remove()}`,
	Dependencies: []*Function{},
}

var WaitIdle = &Function{
	Name:         "waitIdle",
	Definition:   `e=>new Promise(t=>{window.requestIdleCallback(t,{timeout:e})})`,
	Dependencies: []*Function{},
}

var WaitLoad = &Function{
	Name:         "waitLoad",
	Definition:   `function(){const e=this===window;return new Promise((t,n)=>{if(e){if("complete"===document.readyState)return t();window.addEventListener("load",t)}else void 0===this.complete||this.complete?t():(this.addEventListener("load",t),this.addEventListener("error",n))})}`,
	Dependencies: []*Function{},
}

var InputEvent = &Function{
	Name:         "inputEvent",
	Definition:   `function(){this.dispatchEvent(new Event("input",{bubbles:!0})),this.dispatchEvent(new Event("change",{bubbles:!0}))}`,
	Dependencies: []*Function{},
}

var SelectText = &Function{
	Name:         "selectText",
	Definition:   `function(e){const t=this.value.match(new RegExp(e));t&&this.setSelectionRange(t.index,t.index+t[0].length)}`,
	Dependencies: []*Function{},
}

var SelectAllText = &Function{
	Name:         "selectAllText",
	Definition:   `function(){this.select()}`,
	Dependencies: []*Function{},
}

var Select = &Function{
	Name:         "select",
	Definition:   `function(e,t,n){let i;switch(n){case"regex":i=e.map(e=>{const t=new RegExp(e);return e=>t.test(e.innerText)});break;case"css-selector":i=e.map(e=>t=>t.matches(e));break;default:i=e.map(e=>t=>t.innerText.includes(e))}const o=Array.from(this.options);i.forEach(e=>{const n=o.find(e);n&&(n.selected=t)}),this.dispatchEvent(new Event("input",{bubbles:!0})),this.dispatchEvent(new Event("change",{bubbles:!0}))}`,
	Dependencies: []*Function{},
}

var Visible = &Function{
	Name:         "visible",
	Definition:   `function(){const e=functions.tag(this),t=e.getBoundingClientRect(),n=window.getComputedStyle(e);return"none"!==n.display&&"hidden"!==n.visibility&&!!(t.top||t.bottom||t.width||t.height)}`,
	Dependencies: []*Function{Tag},
}

var Invisible = &Function{
	Name:         "invisible",
	Definition:   `function(){return!functions.visible.apply(this)}`,
	Dependencies: []*Function{Visible},
}

var Text = &Function{
	Name:         "text",
	Definition:   `function(){switch(this.tagName){case"INPUT":case"TEXTAREA":return this.value;case"SELECT":return Array.from(this.selectedOptions).map(e=>e.innerText).join();case void 0:return this.textContent;default:return this.innerText}}`,
	Dependencies: []*Function{},
}

var Resource = &Function{
	Name:         "resource",
	Definition:   `function(){return new Promise((e,t)=>{if(this.complete)return e(this.currentSrc);this.addEventListener("load",()=>e(this.currentSrc)),this.addEventListener("error",e=>t(e))})}`,
	Dependencies: []*Function{},
}

var AddScriptTag = &Function{
	Name:         "addScriptTag",
	Definition:   `function(e,t,n){if(!document.getElementById(e))return new Promise((i,o)=>{var s=document.createElement("script");t?(s.src=t,s.onload=i):(s.type="text/javascript",s.text=n,i()),s.id=e,s.onerror=o,document.head.appendChild(s)})}`,
	Dependencies: []*Function{},
}

var AddStyleTag = &Function{
	Name:         "addStyleTag",
	Definition:   `function(e,t,n){if(!document.getElementById(e))return new Promise((i,o)=>{var s;t?((s=document.createElement("link")).rel="stylesheet",s.href=t):((s=document.createElement("style")).type="text/css",s.appendChild(document.createTextNode(n)),i()),s.id=e,s.onload=i,s.onerror=o,document.head.appendChild(s)})}`,
	Dependencies: []*Function{},
}

var FetchAsDataURL = &Function{
	Name:         "fetchAsDataURL",
	Definition:   `e=>fetch(e).then(e=>e.blob()).then(e=>new Promise((t,n)=>{var i=new FileReader;i.onload=(()=>t(i.result)),i.onerror=(()=>n(i.error)),i.readAsDataURL(e)}))`,
	Dependencies: []*Function{},
}

var Selectable = &Function{
	Name:         "selectable",
	Definition:   `e=>e===window?e.document:e`,
	Dependencies: []*Function{},
}

var Tag = &Function{
	Name:         "tag",
	Definition:   `e=>e.tagName?e:e.parentElement`,
	Dependencies: []*Function{},
}
