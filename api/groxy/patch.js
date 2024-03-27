window.alert=console.log;
void async function Patch(){
  if(!globalThis.declare){
    await import(`https://unpkg.com/javaxscript/framework.js?${new Date().getTime()}`);
  }
  if(!select('link[href*="/groxy/injects-css.css"]')){
    body().appendChild(buildElement('link',{attr:{href:"/groxy/injects-css.css",rel:"stylesheet"}}));
  }
  declare(()=>{
    selectApplyAll('img',el=>{
      el.updateAttribute('alt','🐹');
    });
  });
  DOMComplete(()=>{
    declare(()=>{
      selectApplyAll('img:not([natural-width])',el=>{
          el.setAttribute('natural-width',el.naturalWidth);
      });
    });
  });
  declare(()=>{
    queryApplyAll('.syntax-checkbox:not(.active)',el=>el.click());
  });
  
  await import('https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/prism.min.js');
  await import('https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/components/prism-go.min.js');
  importStyle('https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/themes/prism.min.css');
  
  void async function(){
    while(true){
      await sleep(100);
      await nextIdle();
      await nextFrame();
        let code = select(':is(code[class*="language-"], [class*="language-"] code, code[class*="lang-"], [class*="lang-"] code):not([highlighted])');
        if(code){
          Prism.highlightElement(code);
          code.setAttribute('highlighted','on');
        }
    }
  }();
  design(()=>{
    selectApplyAll(':is(html[window-location*="/tour/"] [id="left-side"],html:not([window-location*="/tour/"])) pre:not(.language-go,:has(code))',el=>{
      el.className='language-go';
      el.innerHTML=`<code class="language-go">${el.innerHTML}</code>`;
    });
  });

  design(()=>{
    selectApplyAll(`:is(html[window-location*="/tour/"] [id="left-side"],html:not([window-location*="/tour/"])) code:not(pre>code)`,el=>{
      el.className='language-go';
      el.outerHTML=`<pre class="language-go" style="display:inline-table;margin:0;padding:0;">${el.outerHTML}</pre>`;
    });
  });
  
   design(()=>{
    selectApplyAll(':is(html[window-location*="/tour/"] [id="left-side"],html:not([window-location*="/tour/"])) :is(pre,code):not(.language-go,:has(code.language-go))',el=>{
      el.className='language-go';
    });
  });

  declare(()=>{
    swapText('returns a slice of the string s','returns a string, that is the result of slicing string s,');
  });
  
}();
