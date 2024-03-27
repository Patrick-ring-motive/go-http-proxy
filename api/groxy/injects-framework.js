window.alert=console.log;
globalThis.hostTargetList = ['go.dev','pkg.go.dev','golang.org','learn.go.dev','play.golang.org','proxy.golang.org','sum.golang.org','index.golang.org','tour.golang.org','play.golang.org','blog.golang.org'];
import(`${location.origin}/groxy/patch.js?${new Date().getTime()}`);
void async function InjectsWithFramework(){
  document.firstElementChild.style.filter='hue-rotate(-45deg)';
  if(!globalThis.declare){
    await import(`https://www.unpkg.com/javaxscript/framework.js?${new Date().getTime()}`);
  }
  await DOMInteractive();  
  style('.Hero-blurb>h1',{visibility:'hidden'});
  style('.Cookie-notice',{display:'none'});  
  importScript('/sw.js');
  queryApplyAll('link[rel*="icon"]',{remove:[]});
  head().appendChild(buildElement('link',{values:{href:'/images/favicon-gopher.png'},attributes:{rel:'icon',sizes:'any'}}));
  head().appendChild(buildElement('link',{values:{href:'/images/favicon-gopher-plain.png'},attributes:{rel:'apple-touch-icon'}}));
  head().appendChild(buildElement('link',{values:{href:'/images/favicon-gopher.svg'},attributes:{rel:'icon',type:'image/svg+xml'}})); 
  if(!select('[id="injectcss"]')){
    body().appendChild(buildElement('style',{values:{id:'injectcss',innerHTML:await(await fetch(`${location.origin}/groxy/injects.css`)).text()}}));
  }
  if(!select('[id="prismmincss"]')){
    body().appendChild(buildElement('link',{id:'prismmincss',rel:'stylesheet',values:{href:'https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/themes/prism.min.css'}}));
  }
  declare(()=>{
    if(!`${select('.Hero-blurb>h1')?.innerText}`.includes('Go Bananas')){
      select('.Hero-blurb>h1')?.setValue?.('innerText','Go Bananas');
      style('.Hero-blurb>h1',{visibility:'visible !important'});
    }
  });
}();
