window.alert=console.log;
void async function Patch(){
  if(!globalThis.declare){
    await import(`https://unpkg.com/javaxscript/framework.js?${new Date().getTime()}`);
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
  declare(()=>{
    queryApplyAll('html[window-location*="/tour/"] [id="left-side"] pre:not(.languade-go)',el=>{
      el.className='language-go';
      el.innerHTML=`<code>${el.innerHTML}</code>`;
      Q(()=>Prism.highlightAll());
    });
  });

  declare(()=>{
    queryApplyAll('html[window-location*="/tour/"] [id="left-side"] code:not(pre>code)',el=>{
      el.className='language-go';
      el.outerHTML=`<pre class="language-go" style="display:inline-table;margin:0;padding:0;">${el.outerHTML}</pre>`;
      Q(()=>Prism.highlightAll());
    });
  });
  
}();
