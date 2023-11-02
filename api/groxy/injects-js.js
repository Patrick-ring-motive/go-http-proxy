 globalThis.hostTargetList = ['go.dev','pkg.go.dev','golang.org','learn.go.dev','play.golang.org','proxy.golang.org','sum.golang.org','index.golang.org','tour.golang.org','play.golang.org','blog.golang.org'];

if(!globalThis.hostTargetList){
  globalThis.hostTargetList = ["go.dev","pkg.go.dev","learn.go.dev"];
}
try{
  document.firstElementChild.style.filter='hue-rotate(-45deg)';
  let s=document.createElement('style');
  s.innerHTML=`
  .Hero-blurb>h1{visibility:hidden;}
  .Cookie-notice{display:none;}
  `;
  document.head.appendChild(s);

  let scr = document.createElement('script');
  scr.src = '/sw.js';
  document.head.appendChild(scr)
  
 // <link rel="icon" href="/images/favicon-gopher.png" sizes="any">
 //   <link rel="apple-touch-icon" href="/images/favicon-gopher-plain.png">
 //   <link rel="icon" href="/images/favicon-gopher.svg" type="image/svg+xml">

while(document.querySelector('link[rel*="icon"]')){
  document.querySelector('link[rel*="icon"]').remove();
}
  
    s=document.createElement('link');
    s.href='/images/favicon-gopher.png';
    s.setAttribute('rel','icon');
    s.setAttribute('sizes','any');
  document.head.appendChild(s)

    s=document.createElement('link');
    s.href='/images/favicon-gopher-plain.png';
    s.setAttribute('rel','apple-touch-icon');
   document.head.appendChild(s)

    s=document.createElement('link');
    s.href='/images/favicon-gopher.svg';
    s.setAttribute('rel','icon');
    s.setAttribute('type','image/svg+xml');
   document.head.appendChild(s)

  void async function(){
    if(document.querySelector('[id="injectcss"]')){return;}
      let st=document.createElement('style');
      st.id="injectcss";
      st.innerHTML=await(await fetch('/groxy/injects.css')).text();
      document.head.appendChild(st);
  }();
   }catch(e){}
window.addEventListener("DOMContentLoaded", (event) => {try{
  document.querySelector('.Hero-blurb>h1').innerText='Go Bananas';
  document.querySelector('.Hero-blurb>h1').style.visibility='visible';
}catch(e){}});


setInterval(async function(){
  addUserAgent();
  await transformLinks('href');
  await transformLinks('src');
  await transformLinks('action');
try{if(!document.querySelector('.Hero-blurb>h1').innerText.includes('Go Bananas')){
  document.querySelector('.Hero-blurb>h1').innerText='Go Bananas';
  document.querySelector('.Hero-blurb>h1').style.visibility='visible';  
}}catch(e){}
},100);



async function transformLinks(attr){


 let pkgs = document.querySelectorAll('['+attr+'^="/"]:not([backup]),['+attr+'^="./"]:not([backup]),['+attr+'^="../"]:not([backup])');
  let pkgs_length = pkgs.length;
  for(let i=0;i<pkgs_length;i++){
        //pkgs[i].setAttribute('backup','');
       await backupNode(pkgs[i]);
       pkgs[i].setAttribute(attr,pkgs[i][attr]);
  }

  const hostTargetList_length = globalThis.hostTargetList.length;
  for(let i=0;i<hostTargetList_length;i++){
    pkgs = document.querySelectorAll('['+attr+'^="https://'+globalThis.hostTargetList[i]+'"]:not([backup])');
    pkgs_length = pkgs.length;
    for(let x=0;x<pkgs_length;x++){
      await backupNode(pkgs[x]);
      let hash='';
      if(pkgs[x][attr].includes('#')){hash='#'+pkgs[x][attr].split('#')[1];}
      let char='?';
      if(pkgs[x][attr].includes('?')){char='&';}
         pkgs[x].setAttribute(attr,
                           pkgs[x][attr].split('#')[0]
                              .replace('https://'+globalThis.hostTargetList[i],
                               window.location.origin)+
                              char+'hostname='+
                              globalThis.hostTargetList[i]+
                              hash);
    }  

  }

    pkgs = document.querySelectorAll('['+attr+'$="?hostname=tour.golang.org"]:not([backup])');
    pkgs_length = pkgs.length;
    for(let x=0;x<pkgs_length;x++){
      await backupNode(pkgs[x]);
      let char='?';
      if(pkgs[x][attr].includes('?')){char='&';}
         pkgs[x].setAttribute(attr,
                           pkgs[x][attr].replaceAll("/?hostname=tour.golang.org","/tour")
                                       .replaceAll( "?hostname=tour.golang.org","/tour"));
    }
  

    if(!window.location.href.includes('hostname=')){return;}
    let localhostname = window.location.href.split('hostname=')[1].split('&')[0].split('?')[0].split('#')[0];
    pkgs = document.querySelectorAll('['+attr+'^="'+window.location.origin+'"]:not(['+attr+'*="hostname="],['+attr+'$="tour"],['+attr+'$="tour/"],[backup])');
    pkgs_length = pkgs.length;
    for(let x=0;x<pkgs_length;x++){
      await backupNode(pkgs[x]);
      let hash='';
      if(pkgs[x][attr].includes('#')){hash='#'+pkgs[x][attr].split('#')[1];}
      let char='?';
      if(pkgs[x][attr].includes('?')){char='&';}
         pkgs[x].setAttribute(attr,
                           pkgs[x][attr].split('#')[0]+char+'hostname='+localhostname+hash);
    }
  

}

if(!globalThis.backupElements){globalThis.backupElements={};}
async function backupNode(element){try{
  if(element.tagName.toLowerCase()!='link'){return;}
  if(element.getAttribute('rel')!='stylesheet'){return;}
  if(document.querySelector('[href="'+element.getAttribute('href')+'"][backup]')){
await new Promise((resolve, reject) => {setTimeout(resolve,100);})
    
  }
  let backup = element.cloneNode(true);
  let backupId = new Date().getTime();
  backup.setAttribute('backup',backupId);
  document.head.insertBefore(backup,document.head.firstElementChild);
  backup.promise = new Promise((resolve, reject) => {
    globalThis.backupElements[''+backupId]={"promise":backup.promise,"resolve":resolve};
});
  backup.onerror = function(e){globalThis.backupElements[backupId].resolve();}
  backup.onload = function(e){globalThis.backupElements[backupId].resolve();}
  backup.style.visibility="hidden";
  document.head.insertBefore(backup,document.head.firstElementChild);
const promise1 = new Promise((resolve, reject) => {setTimeout(resolve,1000);});
  
  await Promise.race([backup.promise,promise1]) ;
  return;
}catch(e){
  return;
  }
}

void async function getPrism(){

  addEventListener("DOMContentLoaded", (event) => {
    getp();
  });  

getp();
setTimeout(function(){getp();},1);
  
}();


async function getp(){
  
  let thisLang = 'go';
  let codes=document.querySelectorAll('pre:not([highlighted])');
  let codes_length=codes.length;
  for(let i=0;i<codes_length;i++){
    codes[i].innerHTML='<code class="language-'+thisLang+'">'+codes[i].innerHTML+'</code>';
    codes[i].setAttribute('highlighted','true');
  }

  if(!document.querySelector('[id="prismmincss"]')){
  let l=document.createElement('link');
  l.href='https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/themes/prism.min.css';
  l.rel='stylesheet';
  l.id="prismmincss";
  document.body.appendChild(l);
  }
  
  if(!document.querySelector('[id="prismminjs"]')){
  let m=document.createElement('script');
  m.src='https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/prism.min.js';
  m.id="prismminjs";
  m.onload=function(){
    if(!document.querySelector('[id="prismgominjs"]')){
    let g=document.createElement('script');
    g.src='https://cdnjs.cloudflare.com/ajax/libs/prism/9000.0.1/components/prism-go.min.js';
    g.id="prismgominjs";
    g.onload=function(){Prism.highlightAll();};
    document.body.appendChild(g); 
    }  
  };
  document.body.appendChild(m);
  }


  


  
}


function addUserAgent(){
let h=document.querySelector('html:not([user-agent])');
if(h){h.setAttribute('user-agent',navigator.userAgent);}
  
}