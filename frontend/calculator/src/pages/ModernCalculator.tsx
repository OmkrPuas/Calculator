import React, { useState } from 'react'

type Props = {
  title?: string
  onCalculate?: (n: number | null, err?: string | null) => void
}

export default function ModernCalculator({ title = 'Modern', onCalculate }: Props) {
  const [display, setDisplay] = useState('0')
  const [buffer, setBuffer] = useState<string | null>(null)
  const [operator, setOperator] = useState<string | null>(null)
  const [justComputed, setJustComputed] = useState(false)

  const inputDigit = (d: string) => {
    if (justComputed) {
      setDisplay(d)
      setJustComputed(false)
      return
    }
    setDisplay((cur) => (cur === '0' ? d : cur + d))
  }

  const inputDot = () => {
    if (justComputed) {
      setDisplay('0.')
      setJustComputed(false)
      return
    }
    setDisplay((cur) => (cur.includes('.') ? cur : cur + '.'))
  }

  const clear = () => {
    setDisplay('0')
    setBuffer(null)
    setOperator(null)
    setJustComputed(false)
    if (onCalculate) onCalculate(null, null)
  }

  const chooseOp = (op: string) => {
    if (operator && buffer !== null) {
      const res = compute(Number(buffer), Number(display), operator)
      setDisplay(String(res))
      setBuffer(String(res))
    } else {
      setBuffer(display)
    }
    setDisplay('0')
    setOperator(op)
    setJustComputed(false)
  }

  const compute = (a: number, b: number, op: string) => {
    switch (op) {
      case '+': return a + b
      case '-': return a - b
      case '*': return a * b
      case '/': return b === 0 ? NaN : a / b
      case '^': return Math.pow(a, b)
    }
    return 0
  }

  const equals = () => {
    if (operator && buffer !== null) {
      const res = compute(Number(buffer), Number(display), operator)
      setDisplay(String(res))
      setBuffer(null)
      setOperator(null)
      setJustComputed(true)
      if (onCalculate) onCalculate(Number.isNaN(res) ? null : res, Number.isNaN(res) ? 'division by zero' : null)
    }
  }

  const handleSqrt = () => {
    const v = Number(display)
    if (Number.isNaN(v) || v < 0) {
      setDisplay('ERR')
      if (onCalculate) onCalculate(null, 'invalid input')
      return
    }
    const res = Math.sqrt(v)
    setDisplay(String(res))
    setBuffer(null)
    setOperator(null)
    setJustComputed(true)
    if (onCalculate) onCalculate(res, null)
  }

  const handlePercent = () => {
    const v = Number(display)
    const res = v / 100
    setDisplay(String(res))
    setBuffer(null)
    setOperator(null)
    setJustComputed(true)
    if (onCalculate) onCalculate(res, null)
  }

  return (
    <div className="modern-calc">
      <div className="modern-panel">
        <div className="modern-row">
          <div className="modern-display">{display}</div>
        </div>
        <div className="modern-keys">
          <button type="button" onClick={()=>chooseOp('^')}>^</button>
          <button type="button" onClick={handleSqrt}>√</button>
          <button type="button" onClick={handlePercent}>%</button>
          <button type="button" className="op" onClick={()=>chooseOp('/')}>÷</button>
          <button type="button" onClick={()=>inputDigit('7')}>7</button>
          <button type="button" onClick={()=>inputDigit('8')}>8</button>
          <button type="button" onClick={()=>inputDigit('9')}>9</button>
          <button type="button" className="op" onClick={()=>chooseOp('*')}>×</button>
          <button type="button" onClick={()=>inputDigit('4')}>4</button>
          <button type="button" onClick={()=>inputDigit('5')}>5</button>
          <button type="button" onClick={()=>inputDigit('6')}>6</button>
          <button type="button" className="op" onClick={()=>chooseOp('-')}>-</button>
          <button type="button" onClick={()=>inputDigit('1')}>1</button>
          <button type="button" onClick={()=>inputDigit('2')}>2</button>
          <button type="button" onClick={()=>inputDigit('3')}>3</button>
          <button type="button" className="op" onClick={()=>chooseOp('+')}>+</button>
          <button type="button" className="zero" onClick={()=>inputDigit('0')}>0</button>
          <button type="button" onClick={inputDot}>.</button>
          <button type="button" className="accent" onClick={clear}>AC</button>
          <button type="button" className="equals" onClick={equals}>=</button>
        </div>
      </div>
    </div>
  )
}
