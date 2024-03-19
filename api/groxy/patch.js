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
  declare(()=>{
    selectApplyAll('img:not([natural-width])',el=>{
        el.setAttribute('natural-width',el.naturalWidth);
    });
  });
  declare(()=>{
    queryApplyAll('.syntax-checkbox:not(.active)',el=>el.click());
  });
}();
