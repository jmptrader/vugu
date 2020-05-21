package domrender

// GENERATED FILE, DO NOT EDIT!  See renderer-js-script-maker.go

const jsHelperScript = "(function(){if(window.vuguRender){return;}\nconst opcodeEnd=0\nconst opcodeClearEl=1\nconst opcodeRemoveOtherAttrs=5\nconst opcodeSetAttrStr=6\nconst opcodeSelectMountPoint=7\nconst opcodeMoveToFirstChild=20\nconst opcodeSetElement=21\nconst opcodeSetText=23\nconst opcodeSetComment=24\nconst opcodeMoveToParent=25\nconst opcodeMoveToNextSibling=26\nconst opcodeRemoveOtherEventListeners=27\nconst opcodeSetEventListener=28\nconst opcodeSetInnerHTML=29\nconst opcodeSetCSSTag=30\nconst opcodeRemoveOtherCSSTags=31\nconst opcodeSetJSTag=32\nconst opcodeRemoveOtherJSTags=33\nconst opcodeSetProperty=35\nconst opcodeSelectQuery=36\nconst opcodeBufferInnerHTML=37\nconst opcodeSetAttrNSStr=38\nconst opcodeSetElementNS=39\nconst opcodeCallback=40\nconst opcodeCallbackLastElement=41\nclass Decoder{constructor(dataView,offset){this.dataView=dataView;this.offset=offset||0;return this;}\nreadUint8(){var ret=this.dataView.getUint8(this.offset);this.offset++;return ret;}\nreadRefToString(){var ret=this.dataView.getUint32(this.offset).toString(16).padStart(8,\"0\")+\nthis.dataView.getUint32(this.offset+4).toString(16).padStart(8,\"0\");this.offset+=8;return ret;}\nreadUint32(){var ret=this.dataView.getUint32(this.offset);this.offset+=4;return ret;}\nreadString(){var len=this.dataView.getUint32(this.offset);var ret=utf8decoder.decode(new DataView(this.dataView.buffer,this.dataView.byteOffset+this.offset+4,len));this.offset+=len+4;return ret;}}\nlet utf8decoder=new TextDecoder();window.vuguGetActiveEvent=function(){let state=window.vuguState||{};window.vuguState=state;return state.activeEvent;}\nwindow.vuguGetActiveEventTarget=function(){let state=window.vuguState||{};window.vuguState=state;return state.activeEvent&&state.activeEvent.target;}\nwindow.vuguGetActiveEventCurrentTarget=function(){let state=window.vuguState||{};window.vuguState=state;return state.activeEvent&&state.activeEvent.currentTarget;}\nwindow.vuguActiveEventPreventDefault=function(){let state=window.vuguState||{};window.vuguState=state;if(state.activeEvent&&state.activeEvent.preventDefault){state.activeEvent.preventDefault();}}\nwindow.vuguActiveEventStopPropagation=function(){let state=window.vuguState||{};window.vuguState=state;if(state.activeEvent&&state.activeEvent.stopPropagation){state.activeEvent.stopPropagation();}}\nwindow.vuguSetEventHandler=function(eventHandlerFunc){let state=window.vuguState||{};window.vuguState=state;state.eventHandlerFunc=eventHandlerFunc;}\nwindow.vuguSetCallbackHandler=function(callbackHandlerFunc){let state=window.vuguState||{};window.vuguState=state;state.callbackHandlerFunc=callbackHandlerFunc;}\nwindow.vuguGetRenderArray=function(){if(!window.vuguRenderArray){window.vuguRenderArray=new Uint8Array(16384);}\nreturn window.vuguRenderArray;}\nwindow.vuguRender=function(){let buffer=window.vuguRenderArray;if(!window.vuguRenderArray){throw \"window.vuguRenderArray is not set\";}\nlet state=window.vuguState||{};window.vuguState=state;let textEncoder=new TextEncoder();let bufferView=new DataView(buffer.buffer,buffer.byteOffset,buffer.byteLength);var decoder=new Decoder(bufferView,0);state.mountPointEl=state.mountPointEl||null;state.el=state.el||null;state.bufferedInnerHTML=state.bufferedInnerHTML||null;state.nextElMove=state.nextElMove||null;state.elAttrNames=state.elAttrNames||{};state.eventHandlerMap=state.eventHandlerMap||{};state.elEventKeys=state.elEventKeys||{};instructionLoop:while(true){let opcode=decoder.readUint8();try{switch(opcode){case opcodeEnd:{break instructionLoop;}\ncase opcodeClearEl:{state.el=null;state.nextElMove=null;break;}\ncase opcodeSetProperty:{let el=state.el;if(!el){throw \"opcodeSetProperty: no current reference\";}\nlet propName=decoder.readString();let propValueJSON=decoder.readString();el[propName]=JSON.parse(propValueJSON);break;}\ncase opcodeSelectQuery:{let selector=decoder.readString();state.el=document.querySelector(selector);state.nextElMove=null;break;}\ncase opcodeSetAttrStr:{let el=state.el;if(!el){throw \"opcodeSetAttrStr: no current reference\";}\nlet attrName=decoder.readString();let attrValue=decoder.readString();el.setAttribute(attrName,attrValue);state.elAttrNames[attrName]=true;break;}\ncase opcodeSetAttrNSStr:{let el=state.el;if(!el){throw \"opcodeSetAttrNSStr: no current reference\";}\nlet attrNamespace=decoder.readString();if(attrNamespace==\"\"){attrNamespace=null}\nlet attrName=decoder.readString();let attrValue=decoder.readString();el.setAttributeNS(attrNamespace,attrName,attrValue);state.elAttrNames[attrName]=true;break;}\ncase opcodeSelectMountPoint:{state.elAttrNames={};state.elEventKeys={};let selector=decoder.readString();let nodeName=decoder.readString();if(state.mountPointEl){state.el=state.mountPointEl;}else{let el=document.querySelector(selector);if(!el){throw \"mount point selector not found: \"+selector;}\nstate.mountPointEl=el;state.el=el;}\nlet el=state.el;if(el.nodeName.toUpperCase()!=nodeName.toUpperCase()){let newEl=document.createElement(nodeName);el.parentNode.replaceChild(newEl,el);state.mountPointEl=newEl;el=newEl;}\nstate.el=el;state.nextElMove=null;break;}\ncase opcodeRemoveOtherAttrs:{if(!state.el){throw \"no element selected\";}\nif(state.nextElMove){throw \"cannot call opcodeRemoveOtherAttrs when nextElMove is set\";}\nlet rmAttrNames=[];for(let i=0;i<state.el.attributes.length;i++){if(!state.elAttrNames[state.el.attributes[i].name]){rmAttrNames.push(state.el.attributes[i].name);}}\nfor(let i=0;i<rmAttrNames.length;i++){state.el.attributes.removeNamedItem(rmAttrNames[i]);}\nbreak;}\ncase opcodeMoveToParent:{if(state.nextElMove==\"first_child\"){state.nextElMove=null;}else{let p=state.el.parentNode;let e=state.el;while(e.nextSibling){p.removeChild(e.nextSibling);}\nstate.el=p;state.nextElMove=null;}\nbreak;}\ncase opcodeMoveToFirstChild:{if(state.nextElMove){if(state.nextElMove==\"first_child\"){state.el=state.el.firstChild;if(!state.el){throw \"unable to find state.el.firstChild\";}}else if(state.nextElMove==\"next_sibling\"){state.el=state.el.nextSibling;if(!state.el){throw \"unable to find state.el.nextSibling\";}}\nstate.nextElMove=null;}\nif(!state.el){throw \"must have current selection to use opcodeMoveToFirstChild\";}\nstate.nextElMove=\"first_child\";break;}\ncase opcodeMoveToNextSibling:{if(state.nextElMove){if(state.nextElMove==\"first_child\"){state.el=state.el.firstChild;if(!state.el){throw \"unable to find state.el.firstChild\";}}else if(state.nextElMove==\"next_sibling\"){state.el=state.el.nextSibling;if(!state.el){throw \"unable to find state.el.nextSibling\";}}\nstate.nextElMove=null;}\nif(!state.el){throw \"must have current selection to use opcodeMoveToNextSibling\";}\nstate.nextElMove=\"next_sibling\";break;}\ncase opcodeSetElement:{let nodeName=decoder.readString();state.elAttrNames={};state.elEventKeys={};if(state.nextElMove==\"first_child\"){state.nextElMove=null;let newEl=state.el.firstChild;if(newEl){state.el=newEl;}else{newEl=document.createElement(nodeName);state.el.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove==\"next_sibling\"){state.nextElMove=null;let newEl=state.el.nextSibling;if(newEl){state.el=newEl;}else{newEl=document.createElement(nodeName);state.el.parentNode.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove){throw \"bad state.nextElMove value: \"+state.nextElMove;}\nif(state.el.nodeType!=1||state.el.nodeName.toUpperCase()!=nodeName.toUpperCase()){let newEl=document.createElement(nodeName);state.el.parentNode.replaceChild(newEl,state.el);state.el=newEl;}\nbreak;}\ncase opcodeSetElementNS:{let nodeName=decoder.readString();let namespace=decoder.readString();state.elAttrNames={};state.elEventKeys={};if(state.nextElMove==\"first_child\"){state.nextElMove=null;let newEl=state.el.firstChild;if(newEl){state.el=newEl;}else{newEl=document.createElementNS(namespace,nodeName);state.el.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove==\"next_sibling\"){state.nextElMove=null;let newEl=state.el.nextSibling;if(newEl){state.el=newEl;}else{newEl=document.createElementNS(namespace,nodeName);state.el.parentNode.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove){throw \"bad state.nextElMove value: \"+state.nextElMove;}\nif(state.el.nodeType!=1||state.el.nodeName.toUpperCase()!=nodeName.toUpperCase()){let newEl=document.createElementNS(namespace,nodeName);state.el.parentNode.replaceChild(newEl,state.el);state.el=newEl;}\nbreak;}\ncase opcodeSetText:{let content=decoder.readString();if(state.nextElMove==\"first_child\"){state.nextElMove=null;let newEl=state.el.firstChild;if(newEl){state.el=newEl;}else{let newEl=document.createTextNode(content);state.el.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove==\"next_sibling\"){state.nextElMove=null;let newEl=state.el.nextSibling;if(newEl){state.el=newEl;}else{let newEl=document.createTextNode(content);state.el.parentNode.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove){throw \"bad state.nextElMove value: \"+state.nextElMove;}\nif(state.el.nodeType!=3){let newEl=document.createTextNode(content);state.el.parentNode.replaceChild(newEl,state.el);state.el=newEl;}else{state.el.textContent=content;}\nbreak;}\ncase opcodeSetComment:{let content=decoder.readString();if(state.nextElMove==\"first_child\"){state.nextElMove=null;let newEl=state.el.firstChild;if(newEl){state.el=newEl;}else{let newEl=document.createComment(content);state.el.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove==\"next_sibling\"){state.nextElMove=null;let newEl=state.el.nextSibling;if(newEl){state.el=newEl;}else{let newEl=document.createComment(content);state.el.parentNode.appendChild(newEl);state.el=newEl;break;}}else if(state.nextElMove){throw \"bad state.nextElMove value: \"+state.nextElMove;}\nif(state.el.nodeType!=8){let newEl=document.createComment(content);state.el.parentNode.replaceChild(newEl,state.el);state.el=newEl;}else{state.el.textContent=content;}\nbreak;}\ncase opcodeBufferInnerHTML:{let htmlChunk=decoder.readString();state.bufferedInnerHTML=(state.bufferedInnerHTML||\"\")+htmlChunk\nbreak}\ncase opcodeSetInnerHTML:{let html=decoder.readString();if(!state.el){throw \"opcodeSetInnerHTML must have currently selected element\";}\nif(state.nextElMove){throw \"opcodeSetInnerHTML nextElMove must not be set\";}\nif(state.el.nodeType!=1){throw \"opcodeSetInnerHTML currently selected element expected nodeType 1 but has: \"+state.el.nodeType;}\nstate.el.innerHTML=(state.bufferedInnerHTML||\"\")+html;state.bufferedInnerHTML=null\nbreak;}\ncase opcodeRemoveOtherEventListeners:{let positionID=decoder.readString();let emap=state.eventHandlerMap[positionID]||{};let toBeRemoved=[];for(let k in emap){if(!state.elEventKeys[k]){toBeRemoved.push(k);}}\nfor(let i=0;i<toBeRemoved.length;i++){let k=toBeRemoved[i];let f=emap[k];let kparts=k.split(\"|\");state.el.removeEventListener(kparts[0],f,{capture:!!kparts[1],passive:!!kparts[2]});delete emap[k];}\nif(Object.keys(emap).length==0){delete state.eventHandlerMap[positionID];}else{state.eventHandlerMap[positionID]=emap;}\nbreak;}\ncase opcodeSetEventListener:{let positionID=decoder.readString();let eventType=decoder.readString();let capture=decoder.readUint8();let passive=decoder.readUint8();if(!state.el){throw \"must have state.el set in order to call opcodeSetEventListener\";}\nvar eventKey=eventType+\"|\"+(capture?\"1\":\"0\")+\"|\"+(passive?\"1\":\"0\");state.elEventKeys[eventKey]=true;let emap=state.eventHandlerMap[positionID]||{};let f=emap[eventKey];if(!f){f=function(event){state.activeEvent=event;let eventObj={};for(let i in event){let itype=typeof(event[i]);if((itype==\"boolean\"||itype==\"number\"||itype==\"string\")&&true){eventObj[i]=event[i];}}\nif(event.target){eventObj.target={};let et=event.target;for(let i in et){let itype=typeof(et[i]);if((itype==\"boolean\"||itype==\"number\"||itype==\"string\")&&true){eventObj.target[i]=et[i];}}}\nlet fullJSON=JSON.stringify({position_id:positionID,event_type:eventType,capture:!!capture,passive:!!passive,event_summary:eventObj,});let eventBuffer=state.eventBuffer;if(!eventBuffer){eventBuffer=new Uint8Array(16383);state.eventBuffer=eventBuffer;state.eventBufferView=new DataView(eventBuffer.buffer,eventBuffer.byteOffset,eventBuffer.byteLength);}\nlet encodeResultBuffer=textEncoder.encode(fullJSON);state.eventBuffer.set(encodeResultBuffer,4);state.eventBufferView.setUint32(0,encodeResultBuffer.byteLength-encodeResultBuffer.byteOffset);state.eventHandlerFunc.call(null,eventBuffer);state.activeEvent=null;};emap[eventKey]=f;state.el.removeEventListener(eventType,f,{capture:capture,passive:passive});}\nstate.el.addEventListener(eventType,f,{capture:capture,passive:passive});state.eventHandlerMap[positionID]=emap;break;}\ncase opcodeSetCSSTag:{let elementName=decoder.readString();let textContent=decoder.readString();let attrPairsLen=decoder.readUint8();if(attrPairsLen%2!=0){throw \"attrPairsLen is odd number: \"+attrPairsLen;}\nvar attrMap={};for(let i=0;i<attrPairsLen;i+=2){let key=decoder.readString();let val=decoder.readString();attrMap[key]=val;}\nstate.elCSSTagsSet=state.elCSSTagsSet||[];let thisTagKey=textContent;if(elementName==\"link\"){thisTagKey=attrMap[\"href\"];}\nif(thisTagKey==\"\"){this.console.log(\"element\",elementName,\"ignored due to empty key\");break;}\nlet foundTag=null;this.document.querySelectorAll(elementName).forEach(cssEl=>{let cssElKey;if(elementName==\"style\"){cssElKey=cssEl.textContent;}else{cssElKey=cssEl.href;}\nif(thisTagKey==cssElKey){foundTag=cssEl;}});if(!foundTag){let cTag=this.document.createElement(elementName);for(let k in attrMap){cTag.setAttribute(k,attrMap[k]);}\ncTag.vuguCreated=true;if(textContent){cTag.appendChild(document.createTextNode(textContent))}\nthis.document.head.appendChild(cTag);state.elCSSTagsSet.push(cTag);}else{state.elCSSTagsSet.push(foundTag);}\nbreak;}\ncase opcodeRemoveOtherCSSTags:{state.elCSSTagsSet=state.elCSSTagsSet||[];this.document.querySelectorAll('style,link').forEach(cssEl=>{if(!cssEl.vuguCreated){return;}\nif(state.elCSSTagsSet.findIndex(el=>el==cssEl)>=0){return;}\ncssEl.parentNode.removeChild(cssEl);});state.elCSSTagsSet=null;break;}\ncase opcodeCallbackLastElement:{let callbackID=decoder.readUint32();let el=state.el;if(!el){throw \"opcodeCallbackLastElement: no current reference\";}\nstate.callbackHandlerFunc(callbackID,el);break;}\ncase opcodeCallback:{let callbackID=decoder.readUint32();state.callbackHandlerFunc(callbackID);break;}\ndefault:{console.error(\"found invalid opcode\",opcode);return;}}}catch(e){this.console.log(\"Error during instruction loop. Data opcode=\",opcode,\", state.el=\",state.el,\", state.nextElMove=\",state.nextElMove,\", with error: \",e)\nthrow e;}}}})()"
