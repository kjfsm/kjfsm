function wait(duration: number): Promise<string> {
  return new Promise(resolve => {
    setTimeout(() => resolve(`${duration}ms passed`), duration)
  })
}

wait(1000).then(res => { })