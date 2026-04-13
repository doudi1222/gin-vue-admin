import { createSysError } from '@/api/system/sysError'

function sendErrorTip(errorInfo) {
  // 暂时禁用错误记录功能，避免后端服务未启动时的500错误
  // setTimeout(() => {
  //   const errorData = {
  //     form: errorInfo.type,
  //     info: `${errorInfo.message}\nStack: ${errorInfo.stack}${errorInfo.component ? `\nComponent: ${errorInfo.component.name || 'Unknown'}` : ''}${errorInfo.vueInfo ? `\nVue Info: ${errorInfo.vueInfo}` : ''}${errorInfo.source ? `\nSource: ${errorInfo.source}:${errorInfo.lineno}:${errorInfo.colno}` : ''}`,
  //     level: 'error',
  //     solution: null
  //   }
  //   
  //   createSysError(errorData).catch(apiErr => {
  //     console.error('Failed to create error record:', apiErr)
  //   })
  // }, 0)
}
  
  window.addEventListener('unhandledrejection', (event) => {
    // 暂时禁用错误记录，避免控制台报错
    // sendErrorTip({
    //   type: '前端',
    //   message: `错误信息: ${event.reason}`,
    //   stack: `调用栈: ${event.reason?.stack || '没有调用栈信息'}`,
    // });
  });
