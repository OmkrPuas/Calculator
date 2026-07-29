import React, { useState } from 'react'

type Props = {
  onCalculate?: (n: number | null, err?: string | null) => void
}

export default function TraditionalCalculator({ onCalculate }: Props) {
  const [display, setDisplay] = useState('0')
  const [buffer, setBuffer] = useState<string | null>(null)
  const [operator, setOperator] = useState<string | null>(null)
  const [justComputed, setJustComputed] = useState(false)

  const clear = () => {
    setDisplay('0')
    setBuffer(null)
    setOperator(null)
    if (onCalculate) onCalculate(null, null)
  }

  const inputDigit = (d: string) => {
    if (justComputed) {
      // after result, pressing digit starts a new input
      setDisplay(d)
      setBuffer(null)
      setOperator(null)
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

  const chooseOp = (op: string) => {
    if (justComputed) {
      // continue from computed result
      setBuffer(display)
      setOperator(op)
      setJustComputed(false)
      setDisplay('0')
      return
    }

    if (operator && buffer !== null) {
      const res = compute(Number(buffer), Number(display), operator)
      setBuffer(String(res))
      setDisplay(String(res))
    } else {
      setBuffer(display)
    }
    setOperator(op)
    // prepare for next number
    setDisplay('0')
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
    try {
      const v = Number(display)
      if (v < 0) throw new Error('invalid input')
      const res = Math.sqrt(v)
      setDisplay(String(res))
      setBuffer(null)
      setOperator(null)
      setJustComputed(true)
      if (onCalculate) onCalculate(res, null)
    } catch (e: any) {
      setDisplay('ERR')
      if (onCalculate) onCalculate(null, e?.message || 'error')
    }
  }

  const handlePercent = () => {
    const v = Number(display)
    const res = v / 100
    setDisplay(String(res))
    setJustComputed(true)
    if (onCalculate) onCalculate(res, null)
  }

  return (
    <div className="traditional-calc">
      <div className="trad-screen" role="status">{display}</div>
      <div className="trad-keys">
        <div className="num-keys">
          {['7','8','9','4','5','6','1','2','3','0','.','='].map((k)=> (
            k === '=' ? (
              <button key="equals" type="button" className="equals-trad" onClick={equals}>=</button>
            ) : (
              <button key={k} type="button" onClick={()=> k === '.' ? inputDot() : inputDigit(k)}>{k}</button>
            )
          ))}
        </div>
        <div className="ops-grid">
          <button type="button" onClick={()=>chooseOp('+')}>+</button>
          <button type="button" onClick={()=>chooseOp('-')}>-</button>
          <button type="button" onClick={()=>chooseOp('*')}>*</button>
          <button type="button" onClick={()=>chooseOp('/')}>/</button>
          <button type="button" onClick={()=>chooseOp('^')}>^</button>
          <button type="button" onClick={handleSqrt}>√</button>
          <button type="button" onClick={handlePercent}>%</button>
          <button type="button" onClick={clear}>C</button>
        </div>
      </div>
    </div>
  )
}
