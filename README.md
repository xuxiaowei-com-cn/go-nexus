<div align="center" style="text-align: center;">
    <h1>go-nexus</h1>
    <h3>基于 go 语言的 nexus SDK</h3>
    <a target="_blank" href="https://github.com/996icu/996.ICU/blob/master/LICENSE">
        <img alt="License-Anti" src="https://img.shields.io/badge/License-Anti 996-blue.svg">
    </a>
    <a target="_blank" href="https://996.icu/#/zh_CN">
        <img alt="Link-996" src="https://img.shields.io/badge/Link-996.icu-red.svg">
    </a>
    <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=ZieC6s1WB4njfVbrDHYgoNS8YpT26VtF&jump_from=webapi">
        <img alt="QQ群" src="https://img.shields.io/badge/QQ群-696503132-blue.svg"/>
    </a>
</div>

<p></p>

<div align="center" style="text-align: center;">
    <a target="_blank" href="https://work.weixin.qq.com/gm/75cfc47d6a341047e4b6aca7389bdfa8">
        <img alt="企业微信群" src="static/wechat-work.jpg" height="100"/>
    </a>
</div>

<p></p>

<div align="center" style="text-align: center;">
    基于 go 语言的 nexus SDK
</div>

<p></p>

<div align="center" style="text-align: center;">
  为简化开发工作、提高生产率、解决常见问题而生
</div>

<p></p>

<div align="center" style="text-align: center;">

  <a target="_blank" href="https://space.bilibili.com/198580655">
    <img alt="bilibili 粉丝" src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fapi.spencerwoo.com%2Fsubstats%2F%3Fsource%3Dbilibili%26queryKey%3D198580655&label=bilibili%20fans&query=%24.data.totalSubs&logo=bilibili">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 码龄" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&query=%2F%2Fdiv%5B%40class%3D%27person-code-age%27%5D%5B1%5D%2Fspan%5B1%5D%2Ftext%28%29%5B1%5D&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 粉丝" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&prefix=%E7%B2%89%E4%B8%9D&query=%2F%2Fli%5B4%5D%2Fa%5B1%5D%2Fdiv%5B%40class%3D%27user-profile-statistics-num%27%5D%5B1%5D%2Ftext%28%29%5B1%5D&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 访问" src="https://img.shields.io/badge/dynamic/xml?color=orange&label=CSDN&prefix=%E8%AE%BF%E9%97%AE&query=//span[1]/div[@class='user-profile-statistics-num'][1]/text()[1]&url=https%3A%2F%2Fblog.csdn.net%2Fqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://blog.csdn.net/qq_32596527">
    <img alt="CSDN 博客" src="https://img.shields.io/badge/dynamic/json?color=orange&label=CSDN&prefix=%E5%8D%9A%E5%AE%A2&query=%24.data.blog&suffix=%E7%AF%87&url=https%3A%2F%2Fblog.csdn.net%2Fcommunity%2Fhome-api%2Fv1%2Fget-tab-total%3Fusername%3Dqq_32596527&logo=data:image/x-icon;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAxVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zda/P9qhPz/mKr9/7bC/f/Fz/7/ydL+/8HM/v+tu/3/jaH9/156/P8zV/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/z9h/P+gsP3/8fP+/////////////////////////////////////////////////+ru/v+Zqv3/PV/8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9lgPz/6+/+///////////////////////////////////////////////////////////////////////s7/7/Y378/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aoT8//r6/v///////////////////////v7+/+Po/v/R2f7/y9T+/9rg/v/3+f7////////////////////////////j6P7/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Zm/P/w8/7/////////////////5+v+/4ab/f9AYvz/MVX8/zFV/P8xVfz/MVX8/zVY/P9kf/z/tsP9//39/v////////////T2/v8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/sL79/////////////////87W/v8/Yfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ZYD8//L0/v//////n7D9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/0Bh/P/6+/7////////////v8v7/QmP8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/TWz8/3GJ/P8yVvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/e5L8/////////////////5qr/f8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P+mtv3/////////////////XHn8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7/L/f////////////////87Xfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/ydL+////////////+/v+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P/Ezv7////////////9/f7/M1b8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/7G//f////////////////9HZ/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kqX9/////////////////22H/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P9kf/z/////////////////pbX9/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zRX/P/v8v7////////////s7/7/Nln8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/6Ky/f////////////////+Inf3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/RWb8//f4/v////////////H0/v9Kafz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/PV/8/1Jw/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/kKT9/////////////////9vh/v9DZPz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1Fv/P/m6/7//v7+/3aO/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8zVvz/xM79/////////////////+fr/v9viPz/MVX8/zFV/P8xVfz/MVX8/zRX/P+Emf3/8/X+////////////xc/+/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P87Xfz/ztf+///////////////////////i5/7/sL79/5+w/f+ywP3/6u3+//////////////////////+uvP3/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P83Wvz/sL79//7+/v//////////////////////////////////////////////////////3OL+/0Vl/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/aYP8/9Pb/v//////////////////////////////////////9fb+/5yu/f84W/z/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/1d0/P+Spf3/t8T9/8fR/v/Dzv7/qrn9/3uS/P88Xvz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/MVX8/zFV/P8xVfz/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
  </a>

  <a target="_blank" href="https://www.jetbrains.com/idea">
    <img alt="IntelliJ IDEA" src="https://img.shields.io/static/v1?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAAAAAByaaZbAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAHdElNRQfmBRkBICRBfW8eAAABPklEQVRIx+2UTStEYRiGr/kqZWaKMoyQUsoCWdn4AZKPlZSF5fgPVhYWStlZqPkB7GynWSg/QCmhUFOUGCPNZBrlnNtijmneOYZzykrn2p2e536f577P2wsBPvmQlACQg13Mr4b9CCTpuMunQIdGT8gQRCBZAQRPu0B6aRjsvqKXCRcAJO8lTTf3/GQJKF8Bz5481CfMvEnnxtrRduKhPET7R6GWka+UrBW//8Hej3haqZQFYgNz8VDmYbNNT8iSFDdMM1iSyh3uWHsAesNgVc1D7k4hMeISrN0uAOtAwTYF6Smg2UQUYDYbO8qdjS0Cua9CahsgPd8NleuW3WOFRiTV+nTj8mnL5XbixinVlnEJ7L1vkuzcuJT0ejDufL98UTjZmWyJsqFJvT9aBAT8J5zrrXYFF788xn8gCPDCJ2cr3I1zqSjOAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDIyLTA1LTI1VDAxOjMyOjM2KzAwOjAwH/0yeQAAACV0RVh0ZGF0ZTptb2RpZnkAMjAyMi0wNS0yNVQwMTozMjozNiswMDowMG6gisUAAAAASUVORK5CYII=&message=IntelliJ IDEA">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/xuxiaowei-com-cn/go-nexus?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/xuxiaowei-com-cn/go-nexus?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/xuxiaowei-com-cn/go-nexus?logo=github">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/xuxiaowei-com-cn/go-nexus">
  </a>

  <a target="_blank" href="https://gitee.com/xuxiaowei-com-cn/go-nexus">
    <img alt="码云Gitee stars" src="https://gitee.com/xuxiaowei-com-cn/go-nexus/badge/star.svg?theme=blue">
  </a>

  <a target="_blank" href="https://gitee.com/xuxiaowei-com-cn/go-nexus">
    <img alt="码云Gitee forks" src="https://gitee.com/xuxiaowei-com-cn/go-nexus/badge/fork.svg?theme=blue">
  </a>

  <a target="_blank" href="https://gitlab.com/xuxiaowei-com-cn/go-nexus">
    <img alt="Gitlab stars" src="https://badgen.net/gitlab/stars/xuxiaowei-com-cn/go-nexus?icon=gitlab">
  </a>

  <a target="_blank" href="https://gitlab.com/xuxiaowei-com-cn/go-nexus">
    <img alt="Gitlab forks" src="https://badgen.net/gitlab/forks/xuxiaowei-com-cn/go-nexus?icon=gitlab">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="OSCS Status" src="https://www.oscs1024.com/platform/badge/xuxiaowei-com-cn/go-nexus.svg?size=small">
  </a>

  <a target="_blank" href="https://github.com/xuxiaowei-com-cn/go-nexus">
    <img alt="total lines" src="https://tokei.rs/b1/github/xuxiaowei-com-cn/go-nexus">
  </a>

  <a target="_blank" href="https://www.apache.org/licenses/LICENSE-2.0">
    <img alt="code style" src="https://img.shields.io/badge/license-Apache 2-blue">
  </a>
</div>

## [更新文档](CHANGELOG)

## 使用方式

1. 具体示例请查看对应的测试方法

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
go get -u github.com/PuerkitoBio/goquery
go get -u github.com/google/go-querystring
go get -u github.com/hashicorp/go-cleanhttp
go get -u github.com/hashicorp/go-retryablehttp
go get -u github.com/stretchr/testify
```

### mod

```shell
go mod tidy
```

```shell
go mod download
```

### test

```shell
go test ./... -v
```

## 接口

| Path                                                                | Method | Description                                                                                         | Status |   |
|---------------------------------------------------------------------|--------|-----------------------------------------------------------------------------------------------------|--------|---|
| /v1/assets/{id}                                                     | GET    | Get a single asset                                                                                  | ✅      |   |
| /v1/assets/{id}                                                     | DELETE | Delete a single asset                                                                               | ✅      |   |
| /v1/assets                                                          | GET    | List assets                                                                                         | ✅      |   |
| /v1/azureblobstore/test-connection                                  | POST   | Verify connection using supplied Azure Blob Store settings                                          |        |   |
| /v1/blobstores/{name}/quota-status                                  | GET    | Get quota status for a given blob store                                                             | ✅      |   |
| /v1/blobstores/{name}                                               | DELETE | Delete a blob store by name                                                                         |        |   |
| /v1/blobstores                                                      | GET    | List the blob stores                                                                                | ✅      |   |
| /v1/blobstores/file                                                 | POST   | Create a file blob store                                                                            |        |   |
| /v1/blobstores/file/{name}                                          | GET    | Get a file blob store configuration by name                                                         | ✅      |   |
| /v1/blobstores/file/{name}                                          | PUT    | Update a file blob store configuration by name                                                      |        |   |
| /v1/blobstores/s3/{name}                                            | GET    | Get a S3 blob store configuration by name                                                           |        |   |
| /v1/blobstores/s3/{name}                                            | PUT    | Update an S3 blob store configuration by name                                                       |        |   |
| /v1/blobstores/s3                                                   | POST   | Create an S3 blob store                                                                             |        |   |
| /v1/blobstores/azure/{name}                                         | GET    | Get an Azure blob store configuration by name                                                       |        |   |
| /v1/blobstores/azure/{name}                                         | PUT    | Update an Azure blob store configuration by name                                                    |        |   |
| /v1/blobstores/azure                                                | POST   | Create an Azure blob store                                                                          |        |   |
| /v1/components/{id}                                                 | GET    | Get a single component                                                                              | ✅      |   |
| /v1/components/{id}                                                 | DELETE | Delete a single component                                                                           | ✅      |   |
| /v1/components                                                      | GET    | List components                                                                                     | ✅      |   |
| /v1/components                                                      | POST   | Upload a single component                                                                           |        |   |
| /v1/security/content-selectors                                      | GET    | List content selectors                                                                              |        |   |
| /v1/security/content-selectors                                      | POST   | Create a new content selector                                                                       |        |   |
| /v1/security/content-selectors/{name}                               | GET    | Get a content selector by name                                                                      |        |   |
| /v1/security/content-selectors/{name}                               | PUT    | Update a content selector                                                                           |        |   |
| /v1/security/content-selectors/{name}                               | DELETE | Delete a content selector                                                                           |        |   |
| /v1/email                                                           | GET    | Retrieve the current email configuration                                                            |        |   |
| /v1/email                                                           | PUT    | Set the current email configuration                                                                 |        |   |
| /v1/email                                                           | DELETE | Disable and clear the email configuration                                                           |        |   |
| /v1/email/verify                                                    | POST   | Send a test email to the email address provided in the request body                                 |        |   |
| /v1/formats/{format}/upload-specs                                   | GET    | Get upload field requirements for the desired format                                                |        |   |
| /v1/formats/upload-specs                                            | GET    | Get upload field requirements for each supported format                                             |        |   |
| /v1/lifecycle/bounce                                                | PUT    | Bounce lifecycle phase                                                                              |        |   |
| /v1/lifecycle/phase                                                 | GET    | Get current lifecycle phase                                                                         |        |   |
| /v1/lifecycle/phase                                                 | PUT    | Move to new lifecycle phase                                                                         |        |   |
| /v1/iq/verify-connection                                            | POST   | Verify Sonatype Repository Firewall connection                                                      |        |   |
| /v1/iq                                                              | GET    | Get Sonatype Repository Firewall configuration                                                      |        |   |
| /v1/iq                                                              | PUT    | Update Sonatype Repository Firewall configuration                                                   |        |   |
| /v1/iq/enable                                                       | POST   | Enable Sonatype Repository Firewall                                                                 |        |   |
| /v1/iq/disable                                                      | POST   | Disable Sonatype Repository Firewall                                                                |        |   |
| /v1/system/license                                                  | GET    | Get the current license status.                                                                     |        |   |
| /v1/system/license                                                  | POST   | Upload a new license file.                                                                          |        |   |
| /v1/system/license                                                  | DELETE | Uninstall license if present.                                                                       |        |   |
| /v1/read-only/freeze                                                | POST   | Enable read-only                                                                                    |        |   |
| /v1/read-only/force-release                                         | POST   | Forcibly release read-only                                                                          |        |   |
| /v1/read-only/release                                               | POST   | Release read-only                                                                                   |        |   |
| /v1/read-only                                                       | GET    | Get read-only state                                                                                 |        |   |
| /v1/repositories/{repositoryName}/rebuild-index                     | POST   | Schedule a 'Repair - Rebuild repository search' Task. Hosted or proxy repositories only.            |        |   |
| /v1/repositories/{repositoryName}/invalidate-cache                  | POST   | Invalidate repository cache. Proxy or group repositories only.                                      |        |   |
| /v1/repositories/{repositoryName}                                   | GET    | Get repository details                                                                              |        |   |
| /v1/repositories/{repositoryName}                                   | DELETE | Delete repository of any format                                                                     |        |   |
| /v1/repositorySettings                                              | GET    | List repositories                                                                                   |        |   |
| /v1/repositories                                                    | GET    | List repositories                                                                                   | ✅      |   |
| /v1/repositories/maven/group                                        | POST   | Create Maven group repository                                                                       |        |   |
| /v1/repositories/maven/group/{repositoryName}                       | GET    | Get repository                                                                                      | ✅      |   |
| /v1/repositories/maven/group/{repositoryName}                       | PUT    | Update Maven group repository                                                                       |        |   |
| /v1/repositories/maven/hosted/{repositoryName}                      | GET    | Get repository                                                                                      | ✅      |   |
| /v1/repositories/maven/hosted/{repositoryName}                      | PUT    | Update Maven hosted repository                                                                      |        |   |
| /v1/repositories/maven/hosted                                       | POST   | Create Maven hosted repository                                                                      |        |   |
| /v1/repositories/maven/proxy/{repositoryName}                       | GET    | Get repository                                                                                      | ✅      |   |
| /v1/repositories/maven/proxy/{repositoryName}                       | PUT    | Update Maven proxy repository                                                                       |        |   |
| /v1/repositories/maven/proxy                                        | POST   | Create Maven proxy repository                                                                       |        |   |
| /v1/repositories/apt/hosted/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/apt/hosted/{repositoryName}                        | PUT    | Update APT hosted repository                                                                        |        |   |
| /v1/repositories/apt/hosted                                         | POST   | Create APT hosted repository                                                                        |        |   |
| /v1/repositories/apt/proxy/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/apt/proxy/{repositoryName}                         | PUT    | Update APT proxy repository                                                                         |        |   |
| /v1/repositories/apt/proxy                                          | POST   | Create APT proxy repository                                                                         |        |   |
| /v1/repositories/raw/group                                          | POST   | Create raw group repository                                                                         |        |   |
| /v1/repositories/raw/group/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/raw/group/{repositoryName}                         | PUT    | Update raw group repository                                                                         |        |   |
| /v1/repositories/raw/hosted                                         | POST   | Create raw hosted repository                                                                        |        |   |
| /v1/repositories/raw/hosted/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/raw/hosted/{repositoryName}                        | PUT    | Update raw hosted repository                                                                        |        |   |
| /v1/repositories/raw/proxy                                          | POST   | Create raw proxy repository                                                                         |        |   |
| /v1/repositories/raw/proxy/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/raw/proxy/{repositoryName}                         | PUT    | Update raw proxy repository                                                                         |        |   |
| /v1/repositories/{repositoryName}/health-check                      | POST   | Enable repository health check. Proxy repositories only.                                            |        |   |
| /v1/repositories/{repositoryName}/health-check                      | DELETE | Disable repository health check. Proxy repositories only.                                           |        |   |
| /v1/repositories/npm/group/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/npm/group/{repositoryName}                         | PUT    | Update npm group repository                                                                         |        |   |
| /v1/repositories/npm/group                                          | POST   | Create npm group repository                                                                         |        |   |
| /v1/repositories/npm/hosted                                         | POST   | Create npm hosted repository                                                                        |        |   |
| /v1/repositories/npm/hosted/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/npm/hosted/{repositoryName}                        | PUT    | Update npm hosted repository                                                                        |        |   |
| /v1/repositories/npm/proxy/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/npm/proxy/{repositoryName}                         | PUT    | Update npm proxy repository                                                                         |        |   |
| /v1/repositories/npm/proxy                                          | POST   | Create npm proxy repository                                                                         |        |   |
| /v1/repositories/nuget/group                                        | POST   | Create NuGet group repository                                                                       |        |   |
| /v1/repositories/nuget/group/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/nuget/group/{repositoryName}                       | PUT    | Update NuGet group repository                                                                       |        |   |
| /v1/repositories/nuget/hosted                                       | POST   | Create NuGet hosted repository                                                                      |        |   |
| /v1/repositories/nuget/hosted/{repositoryName}                      | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/nuget/hosted/{repositoryName}                      | PUT    | Update NuGet hosted repository                                                                      |        |   |
| /v1/repositories/nuget/proxy/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/nuget/proxy/{repositoryName}                       | PUT    | Update NuGet proxy repository                                                                       |        |   |
| /v1/repositories/nuget/proxy                                        | POST   | Create NuGet proxy repository                                                                       |        |   |
| /v1/repositories/rubygems/group                                     | POST   | Create RubyGems group repository                                                                    |        |   |
| /v1/repositories/rubygems/group/{repositoryName}                    | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/rubygems/group/{repositoryName}                    | PUT    | Update RubyGems group repository                                                                    |        |   |
| /v1/repositories/rubygems/hosted                                    | POST   | Create RubyGems hosted repository                                                                   |        |   |
| /v1/repositories/rubygems/hosted/{repositoryName}                   | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/rubygems/hosted/{repositoryName}                   | PUT    | Update RubyGems hosted repository                                                                   |        |   |
| /v1/repositories/rubygems/proxy                                     | POST   | Create RubyGems proxy repository                                                                    |        |   |
| /v1/repositories/rubygems/proxy/{repositoryName}                    | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/rubygems/proxy/{repositoryName}                    | PUT    | Update RubyGems proxy repository                                                                    |        |   |
| /v1/repositories/docker/group/{repositoryName}                      | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/docker/group/{repositoryName}                      | PUT    | Update Docker group repository                                                                      |        |   |
| /v1/repositories/docker/group                                       | POST   | Create Docker group repository                                                                      |        |   |
| /v1/repositories/docker/hosted/{repositoryName}                     | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/docker/hosted/{repositoryName}                     | PUT    | Update Docker hosted repository                                                                     |        |   |
| /v1/repositories/docker/hosted                                      | POST   | Create Docker hosted repository                                                                     |        |   |
| /v1/repositories/docker/proxy/{repositoryName}                      | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/docker/proxy/{repositoryName}                      | PUT    | Update Docker group repository                                                                      |        |   |
| /v1/repositories/docker/proxy                                       | POST   | Create Docker proxy repository                                                                      |        |   |
| /v1/repositories/yum/group                                          | POST   | Create Yum group repository                                                                         |        |   |
| /v1/repositories/yum/group/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/yum/group/{repositoryName}                         | PUT    | Update Yum group repository                                                                         |        |   |
| /v1/repositories/yum/hosted/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/yum/hosted/{repositoryName}                        | PUT    | Update Yum hosted repository                                                                        |        |   |
| /v1/repositories/yum/hosted                                         | POST   | Create Yum hosted repository                                                                        |        |   |
| /v1/repositories/yum/proxy                                          | POST   | Create Yum proxy repository                                                                         |        |   |
| /v1/repositories/yum/proxy/{repositoryName}                         | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/yum/proxy/{repositoryName}                         | PUT    | Update Yum proxy repository                                                                         |        |   |
| /v1/repositories/helm/hosted                                        | POST   | Create Helm hosted repository                                                                       |        |   |
| /v1/repositories/helm/hosted/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/helm/hosted/{repositoryName}                       | PUT    | Update Helm hosted repository                                                                       |        |   |
| /v1/repositories/helm/proxy                                         | POST   | Create Helm proxy repository                                                                        |        |   |
| /v1/repositories/helm/proxy/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/helm/proxy/{repositoryName}                        | PUT    | Update Helm proxy repository                                                                        |        |   |
| /v1/repositories/gitlfs/hosted                                      | POST   | Create Git LFS hosted repository                                                                    |        |   |
| /v1/repositories/gitlfs/hosted/{repositoryName}                     | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/gitlfs/hosted/{repositoryName}                     | PUT    | Update Git LFS hosted repository                                                                    |        |   |
| /v1/repositories/pypi/group                                         | POST   | Create PyPI group repository                                                                        |        |   |
| /v1/repositories/pypi/group/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/pypi/group/{repositoryName}                        | PUT    | Update PyPI group repository                                                                        |        |   |
| /v1/repositories/pypi/hosted                                        | POST   | Create PyPI hosted repository                                                                       |        |   |
| /v1/repositories/pypi/hosted/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/pypi/hosted/{repositoryName}                       | PUT    | Update PyPI hosted repository                                                                       |        |   |
| /v1/repositories/pypi/proxy                                         | POST   | Create PyPI proxy repository                                                                        |        |   |
| /v1/repositories/pypi/proxy/{repositoryName}                        | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/pypi/proxy/{repositoryName}                        | PUT    | Update PyPI proxy repository                                                                        |        |   |
| /v1/repositories/conda/proxy                                        | POST   | Create conda proxy repository                                                                       |        |   |
| /v1/repositories/conda/proxy/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/conda/proxy/{repositoryName}                       | PUT    | Update conda proxy repository                                                                       |        |   |
| /v1/repositories/conan/proxy                                        | POST   | Create Conan proxy repository                                                                       |        |   |
| /v1/repositories/conan/proxy/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/conan/proxy/{repositoryName}                       | PUT    | Update Conan proxy repository                                                                       |        |   |
| /v1/repositories/r/group                                            | POST   | Create R group repository                                                                           |        |   |
| /v1/repositories/r/group/{repositoryName}                           | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/r/group/{repositoryName}                           | PUT    | Update R group repository                                                                           |        |   |
| /v1/repositories/r/hosted                                           | POST   | Create R hosted repository                                                                          |        |   |
| /v1/repositories/r/hosted/{repositoryName}                          | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/r/hosted/{repositoryName}                          | PUT    | Update R hosted repository                                                                          |        |   |
| /v1/repositories/r/proxy                                            | POST   | Create R proxy repository                                                                           |        |   |
| /v1/repositories/r/proxy/{repositoryName}                           | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/r/proxy/{repositoryName}                           | PUT    | Update R proxy repository                                                                           |        |   |
| /v1/repositories/cocoapods/proxy                                    | POST   | Create Cocoapods proxy repository                                                                   |        |   |
| /v1/repositories/cocoapods/proxy/{repositoryName}                   | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/cocoapods/proxy/{repositoryName}                   | PUT    | Update Cocoapods proxy repository                                                                   |        |   |
| /v1/repositories/go/group                                           | POST   | Create a Go group repository                                                                        |        |   |
| /v1/repositories/go/group/{repositoryName}                          | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/go/group/{repositoryName}                          | PUT    | Update a Go group repository                                                                        |        |   |
| /v1/repositories/go/proxy                                           | POST   | Create a Go proxy repository                                                                        |        |   |
| /v1/repositories/go/proxy/{repositoryName}                          | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/go/proxy/{repositoryName}                          | PUT    | Update a Go proxy repository                                                                        |        |   |
| /v1/repositories/p2/proxy                                           | POST   | Create p2 proxy repository                                                                          |        |   |
| /v1/repositories/p2/proxy/{repositoryName}                          | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/p2/proxy/{repositoryName}                          | PUT    | Update p2 proxy repository                                                                          |        |   |
| /v1/repositories/bower/group                                        | POST   | Create Bower group repository                                                                       |        |   |
| /v1/repositories/bower/group/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/bower/group/{repositoryName}                       | PUT    | Update Bower group repository                                                                       |        |   |
| /v1/repositories/bower/hosted                                       | POST   | Create Bower hosted repository                                                                      |        |   |
| /v1/repositories/bower/hosted/{repositoryName}                      | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/bower/hosted/{repositoryName}                      | PUT    | Update Bower hosted repository                                                                      |        |   |
| /v1/repositories/bower/proxy/{repositoryName}                       | GET    | Get repository                                                                                      |        |   |
| /v1/repositories/bower/proxy/{repositoryName}                       | PUT    | Update Bower proxy repository                                                                       |        |   |
| /v1/repositories/bower/proxy                                        | POST   | Create Bower proxy repository                                                                       |        |   |
| /v1/routing-rules/{name}                                            | GET    | Get a single routing rule                                                                           |        |   |
| /v1/routing-rules/{name}                                            | PUT    | Update a single routing rule                                                                        |        |   |
| /v1/routing-rules/{name}                                            | DELETE | Delete a single routing rule                                                                        |        |   |
| /v1/routing-rules                                                   | GET    | List routing rules                                                                                  |        |   |
| /v1/routing-rules                                                   | POST   | Create a single routing rule                                                                        |        |   |
| /v1/script                                                          | GET    | List all stored scripts                                                                             |        |   |
| /v1/script                                                          | POST   | Add a new script                                                                                    |        |   |
| /v1/script/{name}                                                   | GET    | Read stored script by name                                                                          |        |   |
| /v1/script/{name}                                                   | PUT    | Update stored script by name                                                                        |        |   |
| /v1/script/{name}                                                   | DELETE | Delete stored script by name                                                                        |        |   |
| /v1/script/{name}/run                                               | POST   | Run stored script by name                                                                           |        |   |
| /v1/search/assets                                                   | GET    | Search assets                                                                                       |        |   |
| /v1/search/assets/download                                          | GET    | Search and download asset                                                                           |        |   |
| /v1/search                                                          | GET    | Search components                                                                                   |        |   |
| /v1/security/user-sources                                           | GET    | Retrieve a list of the available user sources.                                                      |        |   |
| /v1/security/anonymous                                              | GET    | Get Anonymous Access settings                                                                       |        |   |
| /v1/security/anonymous                                              | PUT    | Update Anonymous Access settings                                                                    |        |   |
| /v1/security/ldap                                                   | GET    | List LDAP servers                                                                                   |        |   |
| /v1/security/ldap                                                   | POST   | Create LDAP server                                                                                  |        |   |
| /v1/security/ldap/{name}                                            | GET    | Get LDAP server                                                                                     |        |   |
| /v1/security/ldap/{name}                                            | PUT    | Update LDAP server                                                                                  |        |   |
| /v1/security/ldap/{name}                                            | DELETE | Delete LDAP server                                                                                  |        |   |
| /v1/security/ldap/change-order                                      | POST   | Change LDAP server order                                                                            |        |   |
| /v1/security/privileges                                             | GET    | Retrieve a list of privileges.                                                                      |        |   |
| /v1/security/privileges/{privilegeName}                             | GET    | Retrieve a privilege by name.                                                                       |        |   |
| /v1/security/privileges/{privilegeName}                             | DELETE | Delete a privilege by name.                                                                         |        |   |
| /v1/security/privileges/wildcard                                    | POST   | Create a wildcard type privilege.                                                                   |        |   |
| /v1/security/privileges/application                                 | POST   | Create an application type privilege.                                                               |        |   |
| /v1/security/privileges/wildcard/{privilegeName}                    | PUT    | Update a wildcard type privilege.                                                                   |        |   |
| /v1/security/privileges/application/{privilegeName}                 | PUT    | Update an application type privilege.                                                               |        |   |
| /v1/security/privileges/repository-content-selector                 | POST   | Create a repository content selector type privilege.                                                |        |   |
| /v1/security/privileges/repository-admin                            | POST   | Create a repository admin type privilege.                                                           |        |   |
| /v1/security/privileges/repository-view                             | POST   | Create a repository view type privilege.                                                            |        |   |
| /v1/security/privileges/repository-view/{privilegeName}             | PUT    | Update a repository view type privilege.                                                            |        |   |
| /v1/security/privileges/repository-content-selector/{privilegeName} | PUT    | Update a repository content selector type privilege.                                                |        |   |
| /v1/security/privileges/repository-admin/{privilegeName}            | PUT    | Update a repository admin type privilege.                                                           |        |   |
| /v1/security/privileges/script                                      | POST   | Create a script type privilege.                                                                     |        |   |
| /v1/security/privileges/script/{privilegeName}                      | PUT    | Update a script type privilege.                                                                     |        |   |
| /v1/security/realms/active                                          | GET    | List the active realm IDs in order                                                                  |        |   |
| /v1/security/realms/active                                          | PUT    | Set the active security realms in the order they should be used                                     |        |   |
| /v1/security/realms/available                                       | GET    | List the available realms                                                                           |        |   |
| /v1/security/roles                                                  | GET    | List roles                                                                                          |        |   |
| /v1/security/roles                                                  | POST   | Create role                                                                                         |        |   |
| /v1/security/roles/{id}                                             | GET    | Get role                                                                                            |        |   |
| /v1/security/roles/{id}                                             | PUT    | Update role                                                                                         |        |   |
| /v1/security/roles/{id}                                             | DELETE | Delete role                                                                                         |        |   |
| /v1/security/users/{userId}                                         | PUT    | Update an existing user.                                                                            |        |   |
| /v1/security/users/{userId}                                         | DELETE | Delete a user.                                                                                      |        |   |
| /v1/security/users/{userId}/change-password                         | PUT    | Change a user's password.                                                                           |        |   |
| /v1/security/users                                                  | GET    | Retrieve a list of users. Note if the source is not 'default' the response is limited to 100 users. | ✅      |   |
| /v1/security/users                                                  | POST   | Create a new user in the default source.                                                            |        |   |
| /v1/security/ssl/truststore                                         | GET    | Retrieve a list of certificates added to the trust store.                                           |        |   |
| /v1/security/ssl/truststore                                         | POST   | Add a certificate to the trust store.                                                               |        |   |
| /v1/security/ssl/truststore/{id}                                    | DELETE | Remove a certificate in the trust store.                                                            |        |   |
| /v1/security/ssl                                                    | GET    | Helper method to retrieve certificate details from a remote system.                                 |        |   |
| /v1/status/check                                                    | GET    | Health check endpoint that returns the results of the system status checks                          | ✅      |   |
| /v1/status                                                          | GET    | Health check endpoint that validates server can respond to read requests                            | ✅      |   |
| /v1/status/writable                                                 | GET    | Health check endpoint that validates server can respond to read and write requests                  | ✅      |   |
| /v1/support/supportzip                                              | POST   | Creates and downloads a support zip                                                                 |        |   |
| /v1/support/supportzippath                                          | POST   | Creates a support zip and returns the path                                                          |        |   |
| /v1/tasks                                                           | GET    | List tasks                                                                                          |        |   |
| /v1/tasks/{id}                                                      | GET    | Get a single task by id                                                                             |        |   |
| /v1/tasks/{id}/run                                                  | POST   | Run task                                                                                            |        |   |
| /v1/tasks/{id}/stop                                                 | POST   | Stop task                                                                                           |        |   |
