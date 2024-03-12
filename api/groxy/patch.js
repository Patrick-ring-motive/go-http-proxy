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
    selectApplyAll('img',el=>{
      if(el.naturalWidth==0){
        el.style.display='none';
      }
    });
  });
}();
