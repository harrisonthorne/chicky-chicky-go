$$
{1 \over 2} m_1 {v_1}^2 + {1 \over 2} m_2 {v_2}^2 = \text{sum of initial KE} = s_\text{KE}
$$

$$
m_1 v_1 + m_2 v_2 = \text{sum of initial momentum} = s_p
$$

Solution for $v_1$:

$$
m_1 v_1 + m_2 v_2 = s_{p}
$$

$$
v_1 = {s_p - m_2 v_2 \over m_1}
$$

Solution for $v_2$:

$$
{1 \over 2} m_1 v_1^2 + {1 \over 2} m_2 {v_2}^2 = s_\text{KE}
$$

$$
m_1 {v_1}^2 + m_2 {v_2}^2 = 2s_\text{KE}
$$

$$
{v_2}^2 = {2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

$$
v_2 = \sqrt{2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

Substitute $v_1$:

$$
v_2 = \sqrt{2s_\text{KE} - m_1 {({s_p - m_2 v_2 \over m_1})}^2 \over m_2}
$$

$$
v_2 = \sqrt{2s_\text{KE} - {(s_p - m_2 v_2)^2 \over m_1} \over m_2}
$$

$$
{v_2}^2 = {2s_\text{KE} - {(s_p - m_2 v_2)^2 \over m_1} \over m_2}
$$

$$
m_2 {v_2}^2 = {2s_\text{KE} - {(s_p - m_2 v_2)^2 \over m_1}}
$$

$$
2s_\text{KE} - m_2 {v_2}^2 = {{(s_p - m_2 v_2)^2 \over m_1}}
$$

$$
m_1(2s_\text{KE} - m_2 {v_2}^2) = {(s_p - m_2 v_2)^2}
$$

$$
m_1(2s_\text{KE} - m_2 {v_2}^2) = {s_p}^2 - 2{m_2 v_2 s_p} + {m_2}^2{v_2}^2
$$

$$
2s_\text{KE}m_1 - m_2 {v_2}^2 m_1 = {s_p}^2 - 2{m_2 v_2 s_p} + {m_2}^2{v_2}^2
$$

$$
0 = {s_p}^2 - 2 m_2 s_p v_2 + {m_2}^2 {v_2}^2 - 2 {s_\text{KE}} m_1 + m_2 m_1 {v_2}^2
$$

$$
0 = (m_2 m_1 + {m_2}^2) {v_2}^2 - (2 m_2 s_p) v_2 + ({s_p}^2 - 2 {s_\text{KE}} m_1)
$$

$$
v_2 = {2 m_2 s_p + \sqrt{(-2 m_2 s_p)^2 - 4(m_2 m_1 + {m_2}^2)({s_p}^2 - 2 s_\text{KE} m_1)} \over {2 (m_2 m_1 + {m_2}^2)}}
$$

<!--
Or, substitute $v_2$:

$$
v_1 = {s_p - m_2 \sqrt{2s_\text{KE} - m_1 {v_1}^2 \over m_2} \over m_1}
$$

$$
m_1 v_1 = s_p - m_2 \sqrt{2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

$$
s_p - m_1 v_1 = m_2 \sqrt{2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

$$
{s_p - m_1 v_1 \over m_2} = \sqrt{2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

$$
({s_p - m_1 v_1 \over m_2})^2 = {2s_\text{KE} - m_1 {v_1}^2 \over m_2}
$$

$$
{(s_p - m_1 v_1)^2 \over m_2} = {2s_\text{KE} - m_1 {v_1}^2}
$$

$$
{{s_p}^2 - (2 m_1 s_p) v_1 + {({m_2}^2)v_1}^2 \over m_2} = {2s_\text{KE} - m_1 {v_1}^2}
$$
-->

Simplify $v_2$:

$$
v_2 = {2 m_2 s_p + \sqrt{4 {m_2}^2 {s_p}^2 - 4(m_2 m_1 + {m_2}^2)({s_p}^2 - 2 s_\text{KE} m_1)} \over {2 (m_2 m_1 + {m_2}^2)}}
$$

$$
(m_2 m_1 + {m_2}^2)({s_p}^2 - 2 s_\text{KE} m_1)
$$
$$
= m_2 m_1 {s_p}^2 - 2 s_\text{KE} {m_1}^2 m_2 + {m_2}^2 {s_p}^2 - 2 s_\text{KE} {m_1} {m_2}^2
$$

$$
v_2 = {2 m_2 s_p + 2\sqrt{{m_2}^2 {s_p}^2 - (m_2 m_1 + {m_2}^2)({s_p}^2 - 2 s_\text{KE} m_1)} \over {2 (m_2 m_1 + {m_2}^2)}}
$$

$$
v_2 = {m_2 s_p + \sqrt{{m_2}^2 {s_p}^2 - (m_2 m_1 + {m_2}^2)({s_p}^2 - 2 s_\text{KE} m_1)} \over {m_2 m_1 + {m_2}^2}}
$$

$$
v_2 = {m_2 s_p + \sqrt{2 {m_2}^2 {s_p}^2 - m_1 m_2 {s_p}^2 - 2 s_\text{KE} {m_1}^2 m_2 - 2 s_\text{KE} {m_1} {m_2}^2} \over {m_2 m_1 + {m_2}^2}}
$$

$$
v_2 = {m_2 s_p + \sqrt{{s_p}^2 m_2(2 {m_2} - m_1) - 2 s_\text{KE} m_1 m_2 (m_1 - m_2)} \over {m_2 (m_1 + m_2)}}
$$

$$
\bold{v_2 = {m_2 s_p + \sqrt{m_2 ({s_p}^2(2 m_2 - m_1) - 2 s_\text{KE} m_1 (m_1 - m_2))} \over {m_2 (m_1 + m_2)}}}
$$
