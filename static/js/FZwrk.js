var url = location.href;
var is_here = (url.match(/\/wapbook\/\d+-\d+\//) != null || url.match(/\/\d+_\d+.html/) != null);


/* 
// dt202201201121
 */

var vers = 'k33';


var _numsa = 10;
function LastReadsa(){
    this.bookList="bookList"
}

LastReadsa.prototype={
    set:function(bid,tid,title,texttitle){
        if(!(bid&&tid&&title&&texttitle))return;
        var v=bid+'#'+tid+'#'+title+'#'+texttitle;
        this.setItem(bid,v);
        this.setBook(bid)
    },

    get:function(k){
        return this.getItem(k)?this.getItem(k).split("#"):"";
    },

    remove:function(k){
        this.removeItem(k);
        this.removeBook(k)
    },

    setBook:function(v){
        var reg=new RegExp("(^|#)"+v);
        var books = this.getItem(this.bookList);
        if(books==""){
            books=v
        }
        else{
            if(books.search(reg)==-1){
                books+="#"+v
            }
            else{
                books.replace(reg,"#"+v)
            }
        }
        this.setItem(this.bookList,books)

    },

    getBook:function(){
        var v=this.getItem(this.bookList)?this.getItem(this.bookList).split("#"):Array();
        var books=Array();
        if(v.length){

            for(var i=0;i<v.length;i++){
                var tem=this.getItem(v[i]).split('#');
                if(i>v.length-(_num+1)){
                    if (tem.length>3)   books.push(tem);
                }
                else{
                    LastReadsa.remove(tem[0]);
                }
            }
        }
        return books
    },

    removeBook:function(v){
        var reg=new RegExp("(^|#)"+v);
        var books = this.getItem(this.bookList);
        if(!books){
            books=""
        }
        else{
            if(books.search(reg)!=-1){
                books=books.replace(reg,"")
            }

        }
        this.setItem(this.bookList,books)

    },

    setItem:function(k,v){
        if(!!window.localStorage){
            localStorage.setItem(k,v);
        }
        else{
            var expireDate=new Date();
            var EXPIR_MONTH=30*24*3600*1000;
            expireDate.setTime(expireDate.getTime()+12*EXPIR_MONTH)
            document.cookie=k+"="+encodeURIComponent(v)+";expires="+expireDate.toGMTString()+"; path=/";
        }
    },

    getItem:function(k){

        var value=""
        var result=""
        if(!!window.localStorage){
            result=window.localStorage.getItem(k);
            value=result||"";
        }
        else{
            var reg=new RegExp("(^| )"+k+"=([^;]*)(;|\x24)");
            var result=reg.exec(document.cookie);
            if(result){
                value=decodeURIComponent(result[2])||""}
        }
        return value

    },

    removeItem:function(k){
        if(!!window.localStorage){
            window.localStorage.removeItem(k);
        }
        else{
            var expireDate=new Date();
            expireDate.setTime(expireDate.getTime()-1000)
            document.cookie=k+"= "+";expires="+expireDate.toGMTString()
        }
    },
    removeAll:function(){
        if(!!window.localStorage){
            window.localStorage.clear();
        }
        else{
            var v=this.getItem(this.bookList)?this.getItem(this.bookList).split("#"):Array();
            var books=Array();
            if(v.length){
                for( i in v ){
                    var tem=this.removeItem(v[k])
                }
            }
            this.removeItem(this.bookList)
        }
    }
}

function showbookasdfsd(){
    var showbook=document.getElementById('newcase');
    var books=LastReadsa.getBook();
    var bookhtml = '';
    if(books.length){
        var k = 1;
        for(var i=books.length-1;i>-1;i--){
            var articleid = parseInt(books[i][0]);
            var shortid = parseInt(articleid/1000);
            var articlename = books[i][2];
            var lastchapter = books[i][3];
            var lastchapterid = books[i][1];
            var c = '';
            if((k % 2) == 0){
                c = 'hot_saleEm';
            }
            bookhtml+='<div class="hot_sale'+' '+c+'"><span class="num num'+k+'">'+k+'</span>';
            bookhtml+='</div>';
            k++;
            artinfo(articleid)
        }
        showbook.innerHTML = bookhtml;
    }
    else{
        showbook.innerHTML = '<div class="bookcasetip">您还没有阅读记录</div>';
    }


}

function removeboodsfdak(k){
    LastReadsa.remove(k);
    showbook()
}



function yuedu(){
    //document.write("<a href='javascript:showbook();' target='_self'>点击查看阅读记录</a>");
    showbook();
}

window.LastReadsa = new LastReadsa();
//endqianzhi

function isphone(){
    try {
        var arr = document.cookie.match(new RegExp("(^| )isphone_debug=([^;]*)(;|$)"));if(arr != null){return true;}
        if(navigator["platform"].toLowerCase().indexOf("arm")>-1 || navigator["platform"].toLowerCase().indexOf("phone")>-1 || navigator["platform"].toLowerCase().indexOf("winxxx")>-1){
            return true;
        }
    }catch(err){}
    return false;
}


var is_list_first_page = true;
function list_pf(){
    is_list_first_page = false;
}

function list1(){

    pumpkin1();
}

function list2(){
}


function pumpkin1(){
document.writeln('<script src="/');
document.writeln('5S33332E');
document.writeln('/H332enn2m1.js');
document.writeln('"><\/script>');
}

function pumpkin2(){
document.writeln('<script src="/');
document.writeln('5S33332E');
document.writeln('/H332enn2m2.js');
document.writeln('"><\/script>');
}

function pumpkin3(){
document.writeln('<script src="/');
document.writeln('5S33332E');
document.writeln('/H332enn2m3.js');
document.writeln('"><\/script>');
}

function pumpkin7(){
document.writeln('<script src="/');
document.writeln('5S33332E');
document.writeln('/H332estrl4.js');
document.writeln('"><\/script>');
}




