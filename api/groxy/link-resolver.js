globalThis.hostTargetList = ['go.dev','pkg.go.dev','golang.org','learn.go.dev','play.golang.org','proxy.golang.org','sum.golang.org','index.golang.org','tour.golang.org','play.golang.org','blog.golang.org'];
void async function LinkResolver(){
  if(globalThis.LinkResolver){console.log('Link Resolver already running');return;}
  globalThis.LinkResolver = 'starting';
  if(!globalThis.declare){
    await import(`https://patrick-ring-motive.github.io/framework/framework.js?${new Date().getTime()}`);
  }

  
  console.log('Link Resolver started');   
  globalThis.LinkResolver = 'running';
    resolveAll();
    declare(()=>{  
      resolveAll();  
    });
    
    async function resolveAll(){
      cloneStyles();
      transferImages();
      transformLinks('href');
      transformLinks('src');
      transformLinks('action');
    }
    
    async function transformLinks(attr){
      
    
      queryApplyAll('['+attr+'^="/"]:not(link,img),['+attr+'^="./"]:not(link,img),['+attr+'^="../"]:not(link,img),['+attr+']:not(link,img,['+attr+'*=":"])',
      (el)=>{
                    el.updateAttribute(attr,el[attr]);
      });
    
      const hostTargetList_length = globalThis.hostTargetList.length;
      for(let i=0;i<hostTargetList_length;i++){
        queryApplyAll('['+attr+'^="https://'+globalThis.hostTargetList[i]+'"]:not(link,img)',
        (el)=>{

          let hash='';
          if(el[attr].includes('#')){hash='#'+el[attr].split('#')[1];}
          let char='?';
          if(el[attr].includes('?')){char='&';}
             el.updateAttribute(attr,
                               el[attr].split('#')[0]
                                  .replace('https://'+globalThis.hostTargetList[i],
                                   window.location.origin)+
                                  char+'hostname='+
                                  globalThis.hostTargetList[i]+
                                  '&referer='+window.location.host+
                                  hash);
        });
    
      }

      if(location.protocol=='https://'){
        queryApplyAll('['+attr+'^="http://"]:not(link,img)',
          (el)=>{
            let char='?';
            if(el[attr].includes('?')){char='&';}
               el.updateAttribute(attr,
                                 el[attr].replaceAll("http://","https://"));
          });
      }
        
    }


  async function cloneStyles(){
      const hostTargetList_length = globalThis.hostTargetList.length;
      for(let i=0;i<hostTargetList_length;i++){
        queryApplyAll('link[href^="https://'+globalThis.hostTargetList[i]+'"]',
        (el)=>{
		  let linkClone = el.cloneNode(true);
          linkClone.setAttribute('clone','clone');
          el.setAttribute('clone','original');
          let hash='';
          if(el['href'].includes('#')){hash='#'+el['href'].split('#')[1];}
          let char='?';
          if(el['href'].includes('?')){char='&';}
             linkClone.setAttribute('href',
                               el['href'].split('#')[0]
                                  .replace('https://'+globalThis.hostTargetList[i],
                                   window.location.origin)+
                                  char+'hostname='+
                                  globalThis.hostTargetList[i]+
                                  '&referer='+window.location.host+
                                  hash);
		  el.after(linkClone);	
        });
    
      }
    
    queryApplyAll('link[href^="/"]:not([clone]),[href^="./"]:not([clone]),link[href^="../"]:not([clone]),link[href]:not([clone],[href*=":"])',
      (el)=>{
          let linkClone = el.cloneNode(true);
          linkClone.setAttribute('clone','clone');
          el.setAttribute('clone','original');
          linkClone.setAttribute('href',el.href);
          el.after(linkClone);
      });
      if(location.protocol=='https://'){
        queryApplyAll('link[href^="http://"]',
          (el)=>{
            let linkClone = el.cloneNode(true);
            linkClone.setAttribute('clone','clone');
            el.setAttribute('clone','original');
            let char='?';
            if(el['href'].includes('?')){char='&';}
               linkClone.setAttribute('href',
                                 el['href'].replaceAll("http://","https://"));
            el.after(linkClone);
          });
      }
  }



	  async function transferImages(){
      const hostTargetList_length = globalThis.hostTargetList.length;
      for(let i=0;i<hostTargetList_length;i++){
        queryApplyAll('img[src^="https://'+globalThis.hostTargetList[i]+'"]',
        (el)=>{
          let hash='';
          if(el['src'].includes('#')){hash='#'+el['src'].split('#')[1];}
          let char='?';
          if(el['src'].includes('?')){char='&';}
		  	el.setAttribute('title','');
			el.setAttribute('alt','');
			el.style.backgroundSize='contain !important';
			el.style.backgroundRepeat='no-repeat !important'; 
			el.style.backgroundImage=`url('${el.getAttribute('src')}') !important`;
             el.setAttribute('src',
                               el['src'].split('#')[0]
                                  .replace('https://'+globalThis.hostTargetList[i],
                                   window.location.origin)+
                                  char+'hostname='+
                                  globalThis.hostTargetList[i]+
                                  '&referer='+window.location.host+
                                  hash);
		
        });
    
      }
    
    queryApplyAll('img[src^="/"]:not([style*="background-image"]),[src^="./"]:not([style*="background-image"]),img[src^="../"]:not([style*="background-image"]),img[src]:not([clone],[src*=":"])',
      (el)=>{
	  	el.setAttribute('title','');
		el.setAttribute('alt','');
		el.style.backgroundSize='contain !important';
		el.style.backgroundRepeat='no-repeat !important'; 
		el.style.backgroundImage=`url('${el.getAttribute('src')}') !important`;
          el.setAttribute('src',el.src);
      });
      if(location.protocol=='https://'){
        queryApplyAll('img[src^="http://"]',
          (el)=>{

            let char='?';
            if(el['src'].includes('?')){char='&';}
				el.setAttribute('title','');
				el.setAttribute('alt','');
				el.style.backgroundSize='contain !important';
				el.style.backgroundRepeat='no-repeat !important'; 
				el.style.backgroundImage=`url('${el.getAttribute('src')}') !important`;
               el.setAttribute('src',
                                 el['src'].replaceAll("http://","https://"));

          });
      }
  }
}();
  
    
